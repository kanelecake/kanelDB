package internal

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type CmdAble[V any] interface {
	Get(ctx context.Context, key string) *StringCmd
	Set(ctx context.Context, key string, value V, ttl *time.Duration) *BoolCmd
	Exists(ctx context.Context, keys ...string) *BoolCmd
	Del(ctx context.Context, key string) *BoolCmd
}

type Executor[V any] struct {
	data      map[string]V
	dataMutex sync.RWMutex
	ttl       map[string]int64
	ttlMutex  sync.RWMutex
}

func executor() *Executor[any] {
	return &Executor[any]{
		data: make(map[string]any),
		ttl:  make(map[string]int64),
	}
}

func NewExecutor() CmdAble[any] {
	return executor()
}

func (e *Executor[V]) isExist(key string) bool {
	e.ttlMutex.Lock()
	defer e.ttlMutex.Unlock()

	_, ok := e.data[key]
	if !ok {
		return false
	}

	ttl := e.checkTTL(key)
	if !ttl {
		return false
	}

	return true
}

func (e *Executor[V]) checkTTL(key string) bool {
	e.ttlMutex.Lock()
	defer e.ttlMutex.Unlock()

	val, ok := e.ttl[key]
	if !ok {
		return false
	}

	if time.Now().UnixMilli() > val && val != -1 {
		delete(e.ttl, key)
		return false
	}

	return true
}

// Get is used to get data from database
func (e *Executor[V]) Get(ctx context.Context, key string) *StringCmd {
	e.dataMutex.RLock()
	defer e.dataMutex.RUnlock()

	val, ok := e.data[key]
	if !ok {
		return NewStringCmd(ctx, "", Nil)
	}

	ttl := e.checkTTL(key)
	if !ttl {
		return NewStringCmd(ctx, "", Nil)
	}

	return NewStringCmd(ctx, fmt.Sprintf("%s", val), nil)
}

// Set is used to set data to database
func (e *Executor[V]) Set(ctx context.Context, key string, value V, ttl *time.Duration) *BoolCmd {
	e.dataMutex.Lock()
	defer e.dataMutex.Unlock()

	e.data[key] = value

	if ttl != nil {
		e.ttl[key] = time.Now().UnixMilli() + ttl.Milliseconds()
	} else {
		e.ttl[key] = -1
	}

	return NewBoolCmd(ctx, true, nil)
}

// Exists used to check if keys exist in database
func (e *Executor[V]) Exists(ctx context.Context, keys ...string) *BoolCmd {
	e.dataMutex.RLock()
	defer e.dataMutex.RUnlock()

	cmd := NewBoolCmd(ctx, true, nil)

	for _, key := range keys {
		if !e.isExist(key) {
			cmd = NewBoolCmd(ctx, false, Nil)
			break
		}
	}

	return cmd
}

// Del is used to remove some keys from database
func (e *Executor[V]) Del(ctx context.Context, key string) *BoolCmd {
	e.dataMutex.Lock()
	defer e.dataMutex.Unlock()

	if !e.isExist(key) {
		return NewBoolCmd(ctx, false, Nil)
	}

	delete(e.data, key)
	delete(e.ttl, key)

	return NewBoolCmd(ctx, true, nil)
}
