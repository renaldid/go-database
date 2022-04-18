package siswa

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type siswaRepositoryImpl struct {
	DB *sql.DB
}

func NewSiswaRepository(db *sql.DB) SiswaRepository {
	return &siswaRepositoryImpl{db}
}

func (repository *siswaRepositoryImpl) Insert(ctx context.Context, siswas entity.Siswa) (entity.Siswa, error) {
	script := "INSERT INTO siswa(nama, kelas) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, siswas.Nama, siswas.Kelas)
	if err != nil {
		return siswas, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return siswas, err
	}
	siswas.Id = int32(id)
	return siswas, nil
}

func (repository *siswaRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Siswa, error) {
	script := "SELECT id, nama, kelas FROM siswa WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	siswas := entity.Siswa{}
	if err != nil {
		return siswas, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada datanya
		rows.Scan(&siswas.Id, &siswas.Nama, &siswas.Kelas)
		return siswas, nil
	} else {
		//Jika tidak ada data
		return siswas, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *siswaRepositoryImpl) FindAll(ctx context.Context) ([]entity.Siswa, error) {
	script := "SELECT id, nama, kelas FROM siswa"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var siswa []entity.Siswa
	for rows.Next() {
		siswas := entity.Siswa{}
		rows.Scan(&siswas.Id, &siswas.Nama, &siswas.Kelas)
		siswa = append(siswa, siswas)
	}
	return siswa, nil
}

func (repository *siswaRepositoryImpl) Update(ctx context.Context, id int32, siswas entity.Siswa) (entity.Siswa, error) {
	script := "SELECT id, nama, kelas FROM siswa WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return siswas, err
	}
	if rows.Next() {
		script := "UPDATE siswa SET nama = ?, kelas= ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, siswas.Nama, siswas.Kelas, id)
		if err != nil {
			return siswas, err
		}
		siswas.Id = id
		return siswas, nil
	} else {
		return siswas, errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}

func (repository *siswaRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, nama, kelas FROM siswa WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {
		script := "DELETE FROM siswa WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {
		return "Gagal", errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}
