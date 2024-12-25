package repository

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type PelangganRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) domain.Pelanggan
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Pelanggan
	FindById(ctx context.Context, tx *sql.Tx, pelangganId int) (domain.Pelanggan, error)
	Update(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan) domain.Pelanggan
	Delete(ctx context.Context, tx *sql.Tx, pelanggan domain.Pelanggan)
}
