package siswa

import (
	"context"
	"go_database/entity"
)

type SiswaRepository interface {
	Insert(ctx context.Context, siswas entity.Siswa) (entity.Siswa, error)
	FindById(ctx context.Context, id int32) (entity.Siswa, error)
	FindAll(ctx context.Context) ([]entity.Siswa, error)
	Update(ctx context.Context, id int32, siswas entity.Siswa) (entity.Siswa, error)
	Delete(ctx context.Context, id int32) (string, error)
}
