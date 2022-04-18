package menu

import (
	"context"
	"database/sql"
	"errors"
	"go_database/entity"
	"strconv"
)

type menuRepositoryImpl struct {
	DB *sql.DB
}

func NewMenuRepository(db *sql.DB) MenuRepository {
	return &menuRepositoryImpl{DB: db}
}

func (repository *menuRepositoryImpl) Insert(ctx context.Context, menus entity.Menu) (entity.Menu, error) {
	script := "INSERT INTO menu(nama, harga) VALUES (?,?)"
	result, err := repository.DB.ExecContext(ctx, script, menus.Nama, menus.Harga)
	if err != nil {
		return menus, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		return menus, err
	}
	menus.Id = int32(id)
	return menus, nil
}

func (repository *menuRepositoryImpl) FindById(ctx context.Context, id int32) (entity.Menu, error) {
	script := "SELECT id, nama, harga FROM menu WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	menus := entity.Menu{}
	if err != nil {
		return menus, err
	}
	defer rows.Close()
	if rows.Next() {
		//jika ada datanya
		rows.Scan(&menus.Id, &menus.Nama, &menus.Harga)
		return menus, nil
	} else {
		//Jika tidak ada data
		return menus, errors.New("Id " + strconv.Itoa(int(id)) + "Not Found")
	}
}

func (repository *menuRepositoryImpl) FindAll(ctx context.Context) ([]entity.Menu, error) {
	script := "SELECT id, nama, harga FROM menu"
	rows, err := repository.DB.QueryContext(ctx, script)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var menu []entity.Menu
	for rows.Next() {
		menus := entity.Menu{}
		rows.Scan(&menus.Id, &menus.Nama, &menus.Harga)
		menu = append(menu, menus)
	}
	return menu, nil
}

func (repository *menuRepositoryImpl) Update(ctx context.Context, id int32, menus entity.Menu) (entity.Menu, error) {
	script := "SELECT id, nama, harga FROM menu WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return menus, err
	}
	if rows.Next() {
		script := "UPDATE nama SET nama = ?, harga= ? WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, menus.Nama, menus.Harga, id)
		if err != nil {
			return menus, err
		}
		menus.Id = id
		return menus, nil
	} else {
		return menus, errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}

func (repository *menuRepositoryImpl) Delete(ctx context.Context, id int32) (string, error) {
	script := "SELECT id, nama, harga FROM menu WHERE id = ? LIMIT 1"
	rows, err := repository.DB.QueryContext(ctx, script, id)
	defer rows.Close()
	if err != nil {
		return "Gagal", err
	}
	if rows.Next() {
		script := "DELETE FROM nama WHERE id = ?"
		_, err := repository.DB.ExecContext(ctx, script, id)
		if err != nil {
			return "Id" + strconv.Itoa(int(id)) + "Gagal", err
		}
		return "Id" + strconv.Itoa(int(id)) + "Sukses", nil
	} else {
		return "Gagal", errors.New("Id" + strconv.Itoa(int(id)) + "tidak ada")
	}
}
