package database

import (
	"com.jessegrabowski/go-webapp/business/sampling"
	"context"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/pkger"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/markbates/pkger"
)

type postgres struct{}

var pool *pgxpool.Pool

func (d postgres) init(connectionString string) error {
	var err error
	pool, err = pgxpool.Connect(context.Background(), connectionString)
	return err
}

func (d postgres) migrate(connectionString string) error {
	pkger.Include("/resources/postgres")
	migration, err := migrate.New("pkger:///resources/postgres", connectionString)
	if err != nil {
		return err
	}

	err = migration.Up()
	if err != migrate.ErrNoChange {
		return err
	}

	return nil
}

func (d postgres) CreateEntity(ctx context.Context, entity sampling.Entity) (sampling.Entity, error) {
	return executeInTransaction(ctx, pgx.ReadWrite, entity, func(tx pgx.Tx) (sampling.Entity, error) {
		row := tx.QueryRow(ctx, `
			INSERT INTO entities (display_name) 
			VALUES ($1) 
			RETURNING id, display_name;`, entity.DisplayName)
		err := row.Scan(&entity.Id, &entity.DisplayName)
		return entity, err
	})
}

func (d postgres) ReadEntity(ctx context.Context, id int) (sampling.Entity, error) {
	entity := sampling.Entity{Id: -1}
	return executeInTransaction(ctx, pgx.ReadOnly, entity, func(tx pgx.Tx) (sampling.Entity, error) {
		row := tx.QueryRow(ctx, `
			SELECT id, display_name
			FROM entities
			WHERE id = $1`, id)
		err := row.Scan(&entity.Id, &entity.DisplayName)
		return entity, err
	})
}

func (d postgres) UpdateEntity(ctx context.Context, id int, entity sampling.Entity) (sampling.Entity, error) {
	return entity, nil
}

func (d postgres) DeleteEntity(ctx context.Context, id int) error {
	return nil
}

func (d postgres) CreateSample(ctx context.Context, sample sampling.Sample) (sampling.Sample, error) {
	return sample, nil
}

func (d postgres) ReadSamplesForEntity(ctx context.Context, entityId int, limit int, offset int) ([]sampling.Sample, error) {
	return nil, nil
}

func executeInTransaction[R any](ctx context.Context, accessMode pgx.TxAccessMode, defaultReturn R, operation func(tx pgx.Tx) (R, error)) (R, error) {
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{
		IsoLevel:       pgx.ReadCommitted,
		AccessMode:     accessMode,
		DeferrableMode: pgx.NotDeferrable,
	})
	if err != nil {
		return defaultReturn, err
	}
	defer func() {
		if err != nil {
			_ = tx.Rollback(ctx)
		} else {
			_ = tx.Commit(ctx)
		}
	}()
	return operation(tx)
}
