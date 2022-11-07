package iam

import "context"

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserInterface interface {
	Create(ctx context.Context, user User) (User, error)
}
