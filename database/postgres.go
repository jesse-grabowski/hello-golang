package database

import (
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/markbates/pkger"
)

type Postgres struct{}

var pool *pgxpool.Pool

func (d Postgres) init(connectionString string) error {
	var err error
	pool, err = pgxpool.Connect(context.Background(), connectionString)
	return err
}

func (d Postgres) migrate(connectionString string) error {
	pkger.Include("/resources/postgres")
	migration, err := migrate.New("pkger:///resources/postgres", connectionString)
	if err != nil {
		return err
	}
	return migration.Up()
}
