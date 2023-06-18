package internal

import "time"

type KanelDBInstance struct {
	opt  *Options
	Exec CmdAble[any]
}

func (c *KanelDBInstance) clone() *KanelDBInstance {
	clone := *c
	return &clone
}

// WithTimeout used to set a timeout for request
func (c *KanelDBInstance) WithTimeout(timeout time.Duration) *KanelDBInstance {
	opt := c.opt.clone()
	opt.ReadTimeout = timeout
	opt.WriteTimeout = timeout

	clone := c.clone()
	clone.opt = opt

	return clone
}
