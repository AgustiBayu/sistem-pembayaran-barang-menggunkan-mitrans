package repository

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type OrderItemRepository interface {
	Save(ctx context.Context, tx *sql.Tx, orderItem domain.OrederItem) domain.OrederItem
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.OrederItem, map[int]domain.Pesanan, map[int]domain.Pelanggan, map[int]domain.Produk)
	FindById(ctx context.Context, tx *sql.Tx, orderItemId int) (domain.OrederItem, domain.Pesanan, domain.Pelanggan, domain.Produk, error)
	Update(ctx context.Context, tx *sql.Tx, orderItem domain.OrederItem) domain.OrederItem
	Delete(ctx context.Context, tx *sql.Tx, orderItem domain.OrederItem)
}
