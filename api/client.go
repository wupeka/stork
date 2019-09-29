package api

import "github.com/wupeka/stork/state"

type Client struct {
	st *state.State
}

func (c *Client) State() *state.State {
	return c.st
}
