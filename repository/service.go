package repository

import "github.com/beldeveloper/instafags/model"

// Instagram defines the interface of the Instagram repository.
type Instagram interface {
	Followers(account string) ([]model.User, error)
	Following(account string) ([]model.User, error)
}
