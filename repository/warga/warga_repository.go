package warga

import (
	"context"
	"go_database/entity"
)

type WargaRepository interface {
	Insert(ctx context.Context, wargas entity.Warga) (entity.Warga, error)
	FindById(ctx context.Context, id int32) (entity.Warga, error)
	FindAll(ctx context.Context) ([]entity.Warga, error)
	Update(ctx context.Context, id int32, wargas entity.Warga) (entity.Warga, error)
	Delete(ctx context.Context, id int32) (string, error)
}
