package repository

import (
	"context"
	"database/sql"
	"errors"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type PesananRepositoryImpl struct{}

func NewPesananRepository() PesananRepository {
	return &PesananRepositoryImpl{}
}

func (p *PesananRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan {
	SQL := "INSERT INTO pesanans(pelanggan_id,total_amount,status,created_at) VALUES($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, pesanan.PelangganId, pesanan.TotalAmount, pesanan.Status, pesanan.CreatedAt).Scan(&id)
	helper.PanicIfError(err)
	pesanan.Id = id
	return pesanan
}
func (p *PesananRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pesanan, map[int]domain.Pelanggan) {
	SQL := "select ps.id , ps.pelanggan_id, pl.id ,pl.name ,pl.email ,pl.phone ,pl.address ,ps.total_amount ,ps.status ,ps.created_at from pesanans ps inner join pelanggans pl on ps.pelanggan_id = pl.id "
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var pesanans []domain.Pesanan
	pelangganMap := make(map[int]domain.Pelanggan)
	for rows.Next() {
		pesanan := domain.Pesanan{}
		pelanggan := domain.Pelanggan{}
		err := rows.Scan(&pesanan.Id, &pesanan.PelangganId, &pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address, &pesanan.TotalAmount, &pesanan.Status, &pesanan.CreatedAt)
		helper.PanicIfError(err)
		pesanans = append(pesanans, pesanan)
		pelangganMap[pelanggan.Id] = pelanggan
	}
	return pesanans, pelangganMap
}
func (p *PesananRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, pesananId int) (domain.Pesanan, domain.Pelanggan, error) {
	SQL := "select ps.id , ps.pelanggan_id, pl.id ,pl.name ,pl.email ,pl.phone ,pl.address ,ps.total_amount ,ps.status ,ps.created_at from pesanans ps inner join pelanggans pl on ps.pelanggan_id = pl.id  where ps.id = $1"
	rows, err := tx.QueryContext(ctx, SQL, pesananId)
	helper.PanicIfError(err)
	defer rows.Close()

	pesanan := domain.Pesanan{}
	pelanggan := domain.Pelanggan{}
	if rows.Next() {
		err := rows.Scan(&pesanan.Id, &pesanan.PelangganId, &pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address, &pesanan.TotalAmount, &pesanan.Status, &pesanan.CreatedAt)
		helper.PanicIfError(err)
		return pesanan, pelanggan, nil
	} else {
		return pesanan, pelanggan, errors.New("pesanan id not found")
	}
}
func (p *PesananRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan {
	SQL := "UPDATE pesanans SET pelanggan_id = $1,total_amount = $2,status = $3,created_at = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, pesanan.PelangganId, pesanan.TotalAmount, pesanan.Status, pesanan.CreatedAt, pesanan.Id)
	helper.PanicIfError(err)
	return pesanan
}
func (p *PesananRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) {
	SQL := "DELETE FROM pesanans WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, pesanan.Id)
	helper.PanicIfError(err)
}
