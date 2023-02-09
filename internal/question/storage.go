package question

import "context"

type RepositoryQuestion interface {
	Create(ctx context.Context, user *Question) error
	FindWithProjectId(ctx context.Context, id int) ([]Question, error)
}
