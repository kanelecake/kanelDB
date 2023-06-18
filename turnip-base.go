package turnip

import "time"

type turnipClient struct {
	opt  *Options
	exec CmdAble[any]
}

func (c *turnipClient) clone() *turnipClient {
	clone := *c
	return &clone
}

// WithTimeout used to set a timeout for request
func (c *turnipClient) WithTimeout(timeout time.Duration) *turnipClient {
	opt := c.opt.clone()
	opt.ReadTimeout = timeout
	opt.WriteTimeout = timeout

	clone := c.clone()
	clone.opt = opt

	return clone
}
