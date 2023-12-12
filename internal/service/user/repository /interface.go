package repository

import "context"

type UserRepository interface {
	Register(ctx context.Context)
}
