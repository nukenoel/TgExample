package postgresql

import (
	"TelegramBot/config"
	"TelegramBot/pkg/utils"
	"context"
	"fmt"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"time"
)

type Client interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, cfg *config.Config) (pool *pgxpool.Pool, err error) {
	dsn := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s",
		cfg.StorageConfig.Username, cfg.StorageConfig.Password,
		cfg.StorageConfig.Host, cfg.StorageConfig.Port, cfg.StorageConfig.Database,
	)

	err = utils.DoWithTries(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		pool, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}

		return nil
	}, cfg.StorageConfig.MaxAttempts, cfg.StorageConfig.MaxDelaySecond)
	if err != nil {
		log.Fatalf("all attempts are exceeded unable to connect to postgres, descripiton:%v", err)
	}

	return pool, nil
}
