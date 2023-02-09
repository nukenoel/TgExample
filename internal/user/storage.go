package user

import "context"

type RepositoryUser interface {
	Create(ctx context.Context, user *User) error
	FindOne(ctx context.Context, tgId string) (User, error)
}
