package warga

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestWargaInsert(t *testing.T) {
	wargaRepository := NewWargaRepository(go_database.GetConnection())

	ctx := context.Background()
	warga := entity.Warga{
		Nama: "Usop",
		Umur: 40,
	}

	result, err := wargaRepository.Insert(ctx, warga)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	wargaRepository := NewWargaRepository(go_database.GetConnection())

	warga, err := wargaRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(warga)
}

func TestFindAll(t *testing.T) {
	wargaRepository := NewWargaRepository(go_database.GetConnection())

	warga, err := wargaRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, warga := range warga {
		fmt.Println(warga)
	}
}

func TestUpdate(t *testing.T) {
	wargaRepository := NewWargaRepository(go_database.GetConnection())
	ctx := context.Background()
	warga := entity.Warga{
		Nama: "Nami",
		Umur: 40,
	}
	result, err := wargaRepository.Update(ctx, 1, warga)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	wargaRepository := NewWargaRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := wargaRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
