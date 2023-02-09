package project

import (
	"TelegramBot/pkg/postgresql"
	"context"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(client postgresql.Client) RepositoryProject {
	return &repository{client: client}
}

func (r *repository) Create(ctx context.Context, project *Project) error {
	return nil
}

func (r *repository) FindOne(ctx context.Context, id int) (Project, error) {
	return Project{}, nil
}

func (r *repository) FindAllWhereDivisionId(ctx context.Context, id int) ([]Project, error) {
	return nil, nil
}
