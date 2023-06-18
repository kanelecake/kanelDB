package turnip

import (
	"errors"
	"time"
)

var (
	Nil = errors.New("turnip: nil")
)

type Options struct {
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

// clone used to clone an Options object
func (o *Options) clone() *Options {
	clone := *o
	return &clone
}
