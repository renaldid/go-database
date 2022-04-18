package siswa

import (
	"context"
	"fmt"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestSiswaInsert(t *testing.T) {
	siswaRepository := NewSiswaRepository(go_database.GetConnection())

	ctx := context.Background()
	siswa := entity.Siswa{
		Nama:  "Luffy",
		Kelas: 12,
	}

	result, err := siswaRepository.Insert(ctx, siswa)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	siswaRepository := NewSiswaRepository(go_database.GetConnection())

	siswa, err := siswaRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(siswa)
}

func TestFindAll(t *testing.T) {
	siswaRepository := NewSiswaRepository(go_database.GetConnection())

	siswa, err := siswaRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, siswa := range siswa {
		fmt.Println(siswa)
	}
}

func TestUpdate(t *testing.T) {
	siswaRepository := NewSiswaRepository(go_database.GetConnection())
	ctx := context.Background()
	siswa := entity.Siswa{
		Nama:  "Aldi",
		Kelas: 12,
	}
	result, err := siswaRepository.Update(ctx, 1, siswa)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	siswaRepository := NewSiswaRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := siswaRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
