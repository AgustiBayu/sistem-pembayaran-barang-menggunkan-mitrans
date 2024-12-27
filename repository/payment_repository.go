package repository

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type PaymentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
	FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Payment, map[int]domain.Pesanan, map[int]domain.Pelanggan)
	FindById(ctx context.Context, tx *sql.Tx, paymentId int) (domain.Payment, domain.Pesanan, domain.Pelanggan, error)
	Update(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment
	Delete(ctx context.Context, tx *sql.Tx, payment domain.Payment)
}
