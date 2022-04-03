package sampling

import (
	"context"
)

var dao Dao

func Init(samplingDao Dao) {
	dao = samplingDao
}

func createEntity(ctx context.Context, entity Entity) (Entity, error) {
	return dao.CreateEntity(ctx, entity)
}

func readEntity(ctx context.Context, id int) (Entity, error) {
	return dao.ReadEntity(ctx, id)
}
