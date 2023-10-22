package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/muhwyndhamhp/spill/config"
)

var database *Queries

func GetDB(ctx context.Context) *Queries {
	if database != nil {
		return database
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		config.Get(config.DB_HOST),
		config.Get(config.DB_PORT),
		config.Get(config.DB_USER),
		config.Get(config.DB_NAME),
		config.Get(config.DB_PASSWORD),
	)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		panic(err)
	}
	database = New(pool)

	return database
}
