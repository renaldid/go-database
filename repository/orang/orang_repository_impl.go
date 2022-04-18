package orang

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type orangRepositoryImpl struct {
	DB *sql.DB
}

func NewOrangRepository(db *sql.DB) OrangRepository {
	return &orangRepositoryImpl{DB: db}
}

func (repository *orangRepositoryImpl) Insert(ctx context.Context, orangs entity.Orang) (entity.Orang, error) {
	script := "INSERT INTO orang(nama, pekerjaan) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, orangs.Nama, orangs.Pekerjaan)
	if err != nil {
		return orangs, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return orangs, err
	}
	orangs.Id = int32(id)
	return orangs, nil
}

func (repository *orangRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Orang, error) {
	script := "SELECT id, nama, pekerjaan FROM orang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	orangs := entity.Orang{}
	if err != nil {
		return orangs, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada datanya
		rows.Scan(&orangs.Id, &orangs.Nama, &orangs.Pekerjaan)
		return orangs, nil
	} else {
		//Jika tidak ada data
		return orangs, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *orangRepositoryImpl) FindAll(ctx context.Context) ([]entity.Orang, error) {
	script := "SELECT id, nama, pekerjaan FROM orang"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var orang []entity.Orang
	for rows.Next() {
		orangs := entity.Orang{}
		rows.Scan(&orangs.Id, &orangs.Nama, &orangs.Pekerjaan)
		orang = append(orang, orangs)
	}
	return orang, nil
}

func (repository *orangRepositoryImpl) Update(ctx context.Context, id int32, orangs entity.Orang) (entity.Orang, error) {
	script := "SELECT id, nama, pekerjaan FROM orang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return orangs, err
	}
	if rows.Next() {
		script := "UPDATE nama SET nama = ?, pekerjaan= ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, orangs.Nama, orangs.Pekerjaan, id)
		if err != nil {
			return orangs, err
		}
		orangs.Id = id
		return orangs, nil
	} else {
		return orangs, errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}

func (repository *orangRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, nama, pekerjaan FROM orang WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {
		script := "DELETE FROM orang WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {
		return "Gagal", errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}
