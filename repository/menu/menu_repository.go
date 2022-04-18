package menu

import (
	"context"
	"go_database/entity"
)

type MenuRepository interface {
	Insert(ctx context.Context, menus entity.Menu) (entity.Menu, error)
	FindById(ctx context.Context, id int32) (entity.Menu, error)
	FindAll(ctx context.Context) ([]entity.Menu, error)
	Update(ctx context.Context, id int32, menus entity.Menu) (entity.Menu, error)
	Delete(ctx context.Context, id int32) (string, error)
}
