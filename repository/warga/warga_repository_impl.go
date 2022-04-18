package warga

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type wargaRepositoryImpl struct {
	DB *sql.DB
}

func NewWargaRepository(db *sql.DB) WargaRepository {
	return &wargaRepositoryImpl{DB: db}
}

func (repository *wargaRepositoryImpl) Insert(ctx context.Context, wargas entity.Warga) (entity.Warga, error) {
	script := "INSERT INTO warga(nama, umur) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, wargas.Nama, wargas.Umur)
	if err != nil {
		return wargas, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return wargas, err
	}
	wargas.Id = int32(id)
	return wargas, nil
}

func (repository *wargaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Warga, error) {
	script := "SELECT id, nama, umur FROM warga WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	wargas := entity.Warga{}
	if err != nil {
		return wargas, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada datanya
		rows.Scan(&wargas.Id, &wargas.Nama, &wargas.Umur)
		return wargas, nil
	} else {
		//Jika tidak ada data
		return wargas, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *wargaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Warga, error) {
	script := "SELECT id, nama, umur FROM warga"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var warga []entity.Warga
	for rows.Next() {
		wargas := entity.Warga{}
		rows.Scan(&wargas.Id, &wargas.Nama, &wargas.Umur)
		warga = append(warga, wargas)
	}
	return warga, nil
}

func (repository *wargaRepositoryImpl) Update(ctx context.Context, id int32, wargas entity.Warga) (entity.Warga, error) {
	script := "SELECT id, nama, umur FROM warga WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return wargas, err
	}
	if rows.Next() {
		script := "UPDATE nama SET nama = ?, umur= ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, wargas.Nama, wargas.Umur, id)
		if err != nil {
			return wargas, err
		}
		wargas.Id = id
		return wargas, nil
	} else {
		return wargas, errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}

func (repository *wargaRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, nama, umur FROM warga WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {
		script := "DELETE FROM warga WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {
		return "Gagal", errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}
