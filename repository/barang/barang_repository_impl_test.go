package barang

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestBarangInsert(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	ctx := context.Background()
	barang := entity.Barang{
		Nama:   "Meja",
		Jumlah: 200,
	}

	result, err := barangRepository.Insert(ctx, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindById(context.Background(), 20)
	if err != nil {
		panic(err)
	}
	fmt.Println(barang)
}

func TestFindAll(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())

	barang, err := barangRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}
	for _, barang := range barang {
		fmt.Println(barang)
	}
}

func TestDelete(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())
	ctx := context.Background()
	barang := entity.Barang{
		Nama:   "Lampu",
		Jumlah: 2000,
	}
	result, err := barangRepository.Update(ctx, 1, barang)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestUpdate(t *testing.T) {
	barangRepository := NewBarangRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := barangRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
