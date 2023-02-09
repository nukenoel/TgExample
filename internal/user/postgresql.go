package user

import (
	"TelegramBot/pkg/postgresql"
	"context"
	"errors"
	"fmt"
	"github.com/jackc/pgconn"
	log "github.com/sirupsen/logrus"
	"time"
)

type repository struct {
	client postgresql.Client
}

//NewRepository init new repository
func NewRepository(client postgresql.Client) RepositoryUser {
	return &repository{
		client: client,
	}
}

func (r *repository) Create(ctx context.Context, user *User) error {
	q := `
	INSERT INTO
	    users (tg_id, create_at) 
	values
	       ($1, $2)
	`
	if _, err := r.client.Exec(ctx, q, user.TgId, time.Now()); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			pgErr = err.(*pgconn.PgError)
			newErr := fmt.Errorf(fmt.Sprintf("SQL Error: %s, Detail: %s, Where: %s, Code: %s, SQLState: %s", pgErr.Message, pgErr.Detail, pgErr.Where, pgErr.Code, pgErr.SQLState()))
			log.Error(newErr)
			return newErr
		}
		return err
	}

	return nil
}

func (r *repository) FindOne(ctx context.Context, tgId string) (User, error) {
	var u User
	q := `
	SELECT tgId FROM public.author WHERE tg_id = $1
	`
	err := r.client.QueryRow(ctx, q, tgId).Scan(&u.Id, &u.IsAdmin)
	if err != nil {
		return User{}, err
	}
	return u, nil
}
