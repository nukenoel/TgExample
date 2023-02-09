package division

import "context"

type RepositoryDivision interface {
	Create(ctx context.Context, division *Division) error
	FindOne(ctx context.Context, id int) (Division, error)
	//FindAll find all fields that are active
	FindAll(ctx context.Context) ([]Division, error)
}
