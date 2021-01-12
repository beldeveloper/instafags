package controller

import (
	"github.com/beldeveloper/instafags/model"
	"github.com/beldeveloper/instafags/repository"
)

// NewController creates a new instance of the controller.
func NewController(instagram repository.Instagram, exceptions []string) *Controller {
	c := Controller{i: instagram, e: make(map[string]bool)}
	for _, e := range exceptions {
		c.e[e] = true
	}
	return &c
}

// Controller implements the application top-level functionality.
type Controller struct {
	i repository.Instagram
	e map[string]bool
}

// ListFags returns the list of users who are followed by you but not following you.
func (c Controller) ListFags(account string) ([]model.User, error) {
	following, err := c.i.Following(account)
	if err != nil {
		return nil, err
	}
	followers, err := c.i.Followers(account)
	if err != nil {
		return nil, err
	}
	type state struct {
		fag bool
		u   model.User
	}
	states := make(map[string]*state)
	var count int
	for _, u := range following {
		if c.e[u.Username] {
			continue
		}
		states[u.Username] = &state{
			fag: true,
			u:   u,
		}
		count++
	}
	for _, u := range followers {
		if states[u.Username] != nil {
			states[u.Username].fag = false
			count--
		}
	}
	list := make([]model.User, 0, count)
	for _, s := range states {
		if s.fag {
			list = append(list, s.u)
		}
	}
	return list, nil
}
