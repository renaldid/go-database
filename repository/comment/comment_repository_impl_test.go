package comment

import (
	"context"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"go_database"
	"go_database/entity"
	"testing"
)

func TestCommentInsert(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	ctx := context.Background()
	comment := entity.Comment{
		Email:   "joni@gmail.com",
		Comment: "Joni Belajar db golang",
	}

	result, err := commentRepository.Insert(ctx, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestFindById(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commentRepository.FindById(context.Background(), 37)
	if err != nil {
		panic(err)
	}
	fmt.Println(comment)
}

func TestFindAll(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())

	comment, err := commentRepository.FindAll(context.Background())
	if err != nil {
		panic(err)
	}

	for _, comment := range comment {
		fmt.Println(comment)
	}
}

func TestUpdate(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())
	ctx := context.Background()
	comment := entity.Comment{
		Email:   "purwantorenaldi@gmail.com",
		Comment: "kampret belajar golang",
	}
	result, err := commentRepository.Update(ctx, 1, comment)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}

func TestDelete(t *testing.T) {
	commentRepository := NewCommentRepository(go_database.GetConnection())
	ctx := context.Background()
	result, err := commentRepository.Delete(ctx, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
}
