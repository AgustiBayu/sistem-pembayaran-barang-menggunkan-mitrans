package repository

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type ProdukRepository interface {
	Save(ctx context.Context, tx *sql.Tx, produk domain.Produk) domain.Produk
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Produk
	FindById(ctx context.Context, tx *sql.Tx, produkId int) (domain.Produk, error)
	Update(ctx context.Context, tx *sql.Tx, produk domain.Produk) domain.Produk
	Delete(ctx context.Context, tx *sql.Tx, produk domain.Produk)
}
