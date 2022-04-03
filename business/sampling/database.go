package sampling

import "context"

type Dao interface {
	CreateEntity(ctx context.Context, entity Entity) (Entity, error)
	ReadEntity(ctx context.Context, id int) (Entity, error)
	UpdateEntity(ctx context.Context, id int, entity Entity) (Entity, error)
	DeleteEntity(ctx context.Context, id int) error

	CreateSample(ctx context.Context, sample Sample) (Sample, error)
	ReadSamplesForEntity(ctx context.Context, entityId int, limit int, offset int) ([]Sample, error)
}
