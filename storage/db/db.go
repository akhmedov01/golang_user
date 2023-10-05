package db

import (
	"context"
	"fmt"
	"main/config"
	"main/storage"

	"github.com/jackc/pgx/v4/pgxpool"
)

type store struct {
	db    *pgxpool.Pool
	users *userRepo
}

func NewStorage(ctx context.Context, cfg config.Config) (storage.StoregeI, error) {
	config, err := pgxpool.ParseConfig(
		fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
			cfg.PostgresUser,
			cfg.PostgresPassword,
			cfg.PostgresHost,
			cfg.PostgresPort,
			cfg.PostgresDatabase,
		),
	)

	if err != nil {
		fmt.Println("ParseConfig:", err.Error())
		return nil, err
	}

	config.MaxConns = cfg.PostgresMaxConnections
	pool, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		fmt.Println("ConnectConfig:", err.Error())
		return nil, err
	}
	return &store{
		db: pool,
	}, nil
}

func (s *store) User() storage.UserI {
	if s.users == nil {
		s.users = NewUserRepo(s.db)
	}
	return s.users
}
