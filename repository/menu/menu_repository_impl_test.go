package menu

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestMenuInsert(t *testing.T) {
	menuRepository := NewMenuRepository(go_database.GetConnection())

	ctx := context.Background()
	menu := entity.Menu{
		Nama:  "Rendang",
		Harga: 200000,
	}
	result, err := menuRepository.Insert(ctx, menu)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	menuRepository := NewMenuRepository(go_database.GetConnection())

	menu, err := menuRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(menu)
}

func TestFindAll(t *testing.T) {
	menuRepository := NewMenuRepository(go_database.GetConnection())

	menu, err := menuRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, menu := range menu {
		fmt.Println(menu)
	}
}

func TestUpdate(t *testing.T) {
	menuRepository := NewMenuRepository(go_database.GetConnection())
	ctx := context.Background()
	menu := entity.Menu{
		Nama:  "Rendang",
		Harga: 100000,
	}
	result, err := menuRepository.Update(ctx, 1, menu)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	menuRepository := NewMenuRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := menuRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
