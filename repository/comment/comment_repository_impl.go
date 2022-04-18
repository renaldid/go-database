package comment

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type commentRepositoryImpl struct {
	DB *sql.DB
}

func NewCommentRepository(db *sql.DB) CommentRepository {
	return &commentRepositoryImpl{db}
}

func (repository *commentRepositoryImpl) Insert(ctx context.Context, comment entity.Comment) (entity.Comment, error) {
	script := "INSERT INTO comments(email, comment) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment)
	if err != nil {
		return comment, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return comment, err
	}
	comment.Id = int32(id)
	return comment, nil
}

func (repository *commentRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	comment := entity.Comment{}
	if err != nil {
		return comment, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada datanya
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		return comment, nil
	} else {
		//Jika tidak ada data
		return comment, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *commentRepositoryImpl) FindAll(ctx context.Context) ([]entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []entity.Comment
	for rows.Next() {
		comment := entity.Comment{}
		rows.Scan(&comment.Id, &comment.Email, &comment.Comment)
		comments = append(comments, comment)
	}
	return comments, nil
}

func (repository *commentRepositoryImpl) Update(ctx context.Context, id int32, comment entity.Comment) (entity.Comment, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return comment, err
	}
	if rows.Next() {
		script := "UPDATE comment SET email = ?, comment= ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, comment.Email, comment.Comment, id)
		if err != nil {
			return comment, err
		}
		comment.Id = id
		return comment, nil
	} else {
		return comment, errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}

}

func (repository *commentRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, email, comment FROM comments WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {
		script := "DELETE FROM comment WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {
		return "Gagal", errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}
