package turnip

import (
	"context"
	"time"
)

// Client instance
type Client struct {
	client *turnipClient
	opt    *Options
}

// NewClient returns a client to turnip database
func NewClient() *Client {
	return &Client{
		client: &turnipClient{},
	}
}

// WithTimeout used to set a timeout for client request
func (c *Client) WithTimeout(timeout time.Duration) *Client {
	clone := *c
	clone.client = c.client.WithTimeout(timeout)
	return &clone
}

// +++++++++++++++++++++++++++
// + Commands Implementation +
// +++++++++++++++++++++++++++

// Get is used to get some data from database
func (c *Client) Get(ctx context.Context, key string) *StringCmd {
	return c.client.exec.Get(ctx, key)
}

// Set is used to set some data from database
func (c *Client) Set(ctx context.Context, key string, value any, ttl *time.Duration) *BoolCmd {
	return c.client.exec.Set(ctx, key, value, ttl)
}

// Del is used to delete some data from database
func (c *Client) Del(ctx context.Context, key string) *BoolCmd {
	return c.client.exec.Del(ctx, key)
}

// Exists is used to check if data exists in database
func (c *Client) Exists(ctx context.Context, key ...string) *BoolCmd {
	return c.client.exec.Exists(ctx, key...)
}
