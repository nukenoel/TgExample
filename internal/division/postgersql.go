package division

import (
	"TelegramBot/pkg/postgresql"
	"context"
	log "github.com/sirupsen/logrus"
)

type repository struct {
	client postgresql.Client
}

func NewRepository(Client postgresql.Client) RepositoryDivision {
	return &repository{
		client: Client,
	}
}

func (r *repository) Create(ctx context.Context, division *Division) error {
	return nil
}

func (r *repository) FindOne(ctx context.Context, id int) (Division, error) {
	var d Division
	q := `
	SELECT id, name, active, creat_at FROM public.division WHERE id = $1
	`

	err := r.client.QueryRow(ctx, q, id).Scan(&d.Id, &d.Name, &d.Active, &d.CreatAt)
	if err != nil {
		log.Errorf("error while searching for a division")
		return Division{}, err
	}

	return d, nil
}

func (r *repository) FindAll(ctx context.Context) ([]Division, error) {
	q := `
	SELECT id, name, creat_at FROM public.division WHERE active = true
	`
	rows, err := r.client.Query(ctx, q)
	if err != nil {
		log.Errorf("error while take field in division, descripiton:%v", err)
		return nil, err
	}

	divisions := make([]Division, 0, 5) // 5 it's magic number

	for rows.Next() {
		var d Division
		err = rows.Scan(&d.Id, &d.Name, &d.CreatAt)
		if err != nil {
			log.Errorf("error on getting fields from table division, descripiton:%v", err)
			return nil, err
		}
		divisions = append(divisions, d)
	}
	if err = rows.Err(); err != nil {
		log.Errorf("error while getting all field from division table, description:%v", err)
		return nil, err
	}
	return divisions, nil
}
