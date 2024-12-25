package repository

import (
	"context"
	"database/sql"
	"errors"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type PelangganRepositoryImpl struct{}

func NewPelangganRepository() PelangganRepository {
	return &PelangganRepositoryImpl{}
}

func (p *PelangganRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) domain.Pelanggan {
	SQL := "INSERT INTO pelanggans(name, email, phone, address) VALUES ($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, pelanggan.Name, pelanggan.Email, pelanggan.Phone, pelanggan.Address).Scan(&id)
	helper.PanicIfError(err)
	pelanggan.Id = id
	return pelanggan
}
func (p *PelangganRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Pelanggan {
	SQL := "SELECT id, name, email, phone, address FROM pelanggans"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var pelanggans []domain.Pelanggan
	for rows.Next() {
		pelanggan := domain.Pelanggan{}
		err := rows.Scan(&pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address)
		helper.PanicIfError(err)
		pelanggans = append(pelanggans, pelanggan)
	}
	return pelanggans
}
func (p *PelangganRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pelangganId int) (domain.Pelanggan, error) {
	SQL := "SELECT id, name, email, phone, address FROM pelanggans WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, pelangganId)
	helper.PanicIfError(err)
	defer rows.Close()

	pelanggan := domain.Pelanggan{}
	if rows.Next() {
		err := rows.Scan(&pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address)
		helper.PanicIfError(err)
		return pelanggan, nil
	} else {
		return pelanggan, errors.New("pelanggan id not found")
	}
}
func (p *PelangganRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) domain.Pelanggan {
	SQL := "UPDATE pelanggans SET name = $1, email = $2, phone = $3, address = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, pelanggan.Name, pelanggan.Email, pelanggan.Phone, pelanggan.Address, pelanggan.Id)
	helper.PanicIfError(err)
	return pelanggan
}
func (p *PelangganRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) {
	SQL := "DELETE FROM pelanggans WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, pelanggan.Id)
	helper.PanicIfError(err)
}
