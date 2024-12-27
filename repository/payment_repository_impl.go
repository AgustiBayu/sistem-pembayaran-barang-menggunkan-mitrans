package repository

import (
	"context"
	"database/sql"
	"errors"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type PaymentRepositoryImpl struct{}

func NewPaymentRepository() PaymentRepository {
	return &PaymentRepositoryImpl{}
}

func (p *PaymentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "INSERT INTO payments(pesanan_id,payment_method,payment_status,payment_amount,payment_date) VALUES($1,$2,$3,$4,$5) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, payment.PesananId, payment.PaymentMethod, payment.PaymentStatus, payment.PaymentAmount, payment.PaymentDate).Scan(&id)
	helper.PanicIfError(err)
	payment.Id = id
	return payment
}
func (p *PaymentRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.Payment, map[int]domain.Pesanan, map[int]domain.Pelanggan) {
	SQL := "select py.id ,py.pesanan_id, ps.id , ps.pelanggan_id, pl.id ,pl.name ,pl.email ,pl.phone ,pl.address ,ps.total_amount ,ps.status ,ps.created_at,py.payment_method ,py.payment_status ,py.payment_amount ,py.payment_date from payments py inner join pesanans ps on py.pesanan_id = ps.id inner join pelanggans pl on ps.pelanggan_id = pl.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var payments []domain.Payment
	pesananMap := make(map[int]domain.Pesanan)
	pelangganMap := make(map[int]domain.Pelanggan)
	for rows.Next() {
		payment := domain.Payment{}
		pesanan := domain.Pesanan{}
		pelanggan := domain.Pelanggan{}
		err := rows.Scan(&payment.Id, &payment.PesananId, &pesanan.Id, &pesanan.PelangganId, &pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address, &pesanan.TotalAmount, &pesanan.Status, &pesanan.CreatedAt, &payment.PaymentMethod, &payment.PaymentStatus, &payment.PaymentAmount, &payment.PaymentDate)
		helper.PanicIfError(err)
		payments = append(payments, payment)
		pesananMap[pesanan.Id] = pesanan
		pelangganMap[pelanggan.Id] = pelanggan
	}
	return payments, pesananMap, pelangganMap
}
func (p *PaymentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, paymentId int) (domain.Payment, domain.Pesanan, domain.Pelanggan, error) {
	SQL := "select py.id ,py.pesanan_id, ps.id , ps.pelanggan_id, pl.id ,pl.name ,pl.email ,pl.phone ,pl.address ,ps.total_amount ,ps.status ,ps.created_at,py.payment_method ,py.payment_status ,py.payment_amount ,py.payment_date from payments py inner join pesanans ps on py.pesanan_id = ps.id inner join pelanggans pl on ps.pelanggan_id = pl.id where py.id = $1"
	rows, err := tx.QueryContext(ctx, SQL, paymentId)
	helper.PanicIfError(err)
	defer rows.Close()

	payment := domain.Payment{}
	pesanan := domain.Pesanan{}
	pelanggan := domain.Pelanggan{}
	if rows.Next() {
		err := rows.Scan(&payment.Id, &payment.PesananId, &pesanan.Id, &pesanan.PelangganId, &pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address, &pesanan.TotalAmount, &pesanan.Status, &pesanan.CreatedAt, &payment.PaymentMethod, &payment.PaymentStatus, &payment.PaymentAmount, &payment.PaymentDate)
		helper.PanicIfError(err)
		return payment, pesanan, pelanggan, nil
	} else {
		return payment, pesanan, pelanggan, errors.New("payment id not found")
	}
}
func (p *PaymentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, payment domain.Payment) domain.Payment {
	SQL := "UPDATE payments SET pesanan_id = $1,payment_method = $2,payment_status = $3,payment_amount = $4,payment_date = $5 WHERE id = $6"
	_, err := tx.ExecContext(ctx, SQL, payment.PesananId, payment.PaymentMethod, payment.PaymentStatus, payment.PaymentAmount, payment.PaymentDate, payment.Id)
	helper.PanicIfError(err)
	return payment
}
func (p *PaymentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, payment domain.Payment) {
	SQL := "DELETE FROM payments WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, payment.Id)
	helper.PanicIfError(err)
}
