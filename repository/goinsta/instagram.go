package goinsta

import (
	"fmt"

	"github.com/ahmdrz/goinsta/v2"
	"github.com/beldeveloper/instafags/model"
)

// Instagram repository based on goinsta library.
type Instagram struct {
	Client *goinsta.Instagram
}

// Followers returns the list of users who follow you.
func (i Instagram) Followers(account string) ([]model.User, error) {
	acc, err := i.Client.Profiles.ByName(account)
	if err != nil {
		return nil, fmt.Errorf("error while getting account by name: %v", err)
	}
	acc.SetInstagram(i.Client)
	return i.convertUsers(acc.Followers()), nil
}

// Following returns the list of users who are followed by you.
func (i Instagram) Following(account string) ([]model.User, error) {
	acc, err := i.Client.Profiles.ByName(account)
	if err != nil {
		return nil, fmt.Errorf("error while getting account by name: %v", err)
	}
	acc.SetInstagram(i.Client)
	return i.convertUsers(acc.Following()), nil
}

func (i Instagram) convertUsers(users *goinsta.Users) []model.User {
	res := make([]model.User, 0)
	users.SetInstagram(i.Client)
	for users.Next() {
		for _, u := range users.Users {
			res = append(res, model.User{Username: u.Username})
		}
	}
	return res
}
