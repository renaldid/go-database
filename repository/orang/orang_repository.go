package orang

import (
	"context"
	"go_database/entity"
)

type OrangRepository interface {
	Insert(ctx context.Context, orangs entity.Orang) (entity.Orang, error)
	FindById(ctx context.Context, id int32) (entity.Orang, error)
	FindAll(ctx context.Context) ([]entity.Orang, error)
	Update(ctx context.Context, id int32, orangs entity.Orang) (entity.Orang, error)
	Delete(ctx context.Context, id int32) (string, error)
}
