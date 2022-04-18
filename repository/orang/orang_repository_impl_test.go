package orang

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestOrangInsert(t *testing.T) {
	orangRepository := NewOrangRepository(go_database.GetConnection())

	ctx := context.Background()
	orang := entity.Orang{
		Nama:      "Renaldi",
		Pekerjaan: "Progammer Golang",
	}

	result, err := orangRepository.Insert(ctx, orang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	orangRepository := NewOrangRepository(go_database.GetConnection())

	orang, err := orangRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(orang)
}

func TestFindAll(t *testing.T) {
	orangRepository := NewOrangRepository(go_database.GetConnection())

	orang, err := orangRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, orang := range orang {
		fmt.Println(orang)
	}
}

func TestUpdate(t *testing.T) {
	orangRepository := NewOrangRepository(go_database.GetConnection())
	ctx := context.Background()
	orang := entity.Orang{
		Nama:      "Aldi",
		Pekerjaan: "Progammer",
	}
	result, err := orangRepository.Update(ctx, 1, orang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	orangRepository := NewOrangRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := orangRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
