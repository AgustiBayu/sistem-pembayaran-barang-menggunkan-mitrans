package repository

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type PesananRepository interface {
	Save(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Pesanan, map[int]domain.Pelanggan)
	FindById(ctx context.Context, tx *sql.Tx, pesananId int) (domain.Pesanan, domain.Pelanggan, error)
	Update(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan) domain.Pesanan
	Delete(ctx context.Context, tx *sql.Tx, pesanan domain.Pesanan)
}
