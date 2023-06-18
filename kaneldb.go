package kv_database

import (
	"context"
	"github.com/kanelecake/kanelDB/internal"
	"time"
)

// KanelDB instance
type KanelDB struct {
	client *internal.KanelDBInstance
	opt    *internal.Options
}

// NewClient returns a client to turnip database
func NewClient() *KanelDB {
	return &KanelDB{
		client: &internal.KanelDBInstance{
			Exec: internal.NewExecutor(),
		},
	}
}

// WithTimeout used to set a timeout for client request
func (c *KanelDB) WithTimeout(timeout time.Duration) *KanelDB {
	clone := *c
	clone.client = c.client.WithTimeout(timeout)
	return &clone
}

// +++++++++++++++++++++++++++
// + Commands Implementation +
// +++++++++++++++++++++++++++

// Get is used to get some data from database
func (c *KanelDB) Get(ctx context.Context, key string) *internal.StringCmd {
	return c.client.Exec.Get(ctx, key)
}

// Set is used to set some data from database
func (c *KanelDB) Set(ctx context.Context, key string, value any, ttl *time.Duration) *internal.BoolCmd {
	return c.client.Exec.Set(ctx, key, value, ttl)
}

// Del is used to delete some data from database
func (c *KanelDB) Del(ctx context.Context, key string) *internal.BoolCmd {
	return c.client.Exec.Del(ctx, key)
}

// Exists is used to check if data exists in database
func (c *KanelDB) Exists(ctx context.Context, key ...string) *internal.BoolCmd {
	return c.client.Exec.Exists(ctx, key...)
}
