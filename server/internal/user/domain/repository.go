package domain

import "context"

type UserRepository interface {
	FindByEmail(ctx context.Context, email string) (*User, error)
}
