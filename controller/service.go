package controller

import "github.com/beldeveloper/instafags/model"

// Service defines the interface of the application controller.
type Service interface {
	ListFags(account string) ([]model.User, error)
}
