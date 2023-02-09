package project

import "context"

type RepositoryProject interface {
	Create(ctx context.Context, project *Project) error
	FindOne(ctx context.Context, id int) (Project, error)
	FindAllWhereDivisionId(ctx context.Context, id int) ([]Project, error)
}
