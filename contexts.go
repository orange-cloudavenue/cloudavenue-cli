package main

import (
	"encoding/json"
	"time"

	"github.com/zalando/go-keyring"
)

type contexts struct {
	Current string         `json:"current"`
	Items   []*contextInfo `json:"items"`
}

type contextInfo struct {
	ID           string
	Name         string
	Organization string
	Username     string
	Password     string
	CreatedAt    time.Time
	LastUsedAt   time.Time
}

func Contexts() (*contexts, error) {
	ctxs, err := keyring.Get("cav-cli", "contexts")
	if err != nil {
		return nil, err
	}

	contexts := new(contexts)
	err = json.Unmarshal([]byte(ctxs), contexts)
	if err != nil {
		return nil, err
	}

	return contexts, nil
}

func (c *contexts) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	return keyring.Set("cav-cli", "contexts", string(data))
}

func (c *contexts) Add(ctx *contextInfo) {
	// Check if context already exists
	for i, existing := range c.Items {
		if existing.Organization == ctx.Organization {
			// Update existing context
			c.Items[i] = ctx
			return
		}
	}

	c.Items = append(c.Items, ctx)
}

func (c *contexts) Remove(name string) {
	for i, existing := range c.Items {
		if existing.Name == name {
			c.Items = append(c.Items[:i], c.Items[i+1:]...)
			// If the removed context was the current one, unset current
			if c.Current == name {
				c.Current = ""
			}
			return
		}
	}
}

func (c *contexts) SetCurrent(name string) {
	for _, existing := range c.Items {
		if existing.Name == name {
			c.Current = name
			existing.LastUsedAt = time.Now()
			return
		}
	}
}

func (c *contexts) GetCurrent() *contextInfo {
	for _, existing := range c.Items {
		if existing.Name == c.Current {
			return existing
		}
	}
	return nil
}
