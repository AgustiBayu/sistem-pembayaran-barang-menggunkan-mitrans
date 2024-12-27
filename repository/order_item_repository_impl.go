package repository

import (
	"context"
	"database/sql"
	"errors"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type OrderItemRepositoryImpl struct{}

func NewOrderItemRepository() OrderItemRepository {
	return &OrderItemRepositoryImpl{}
}

func (o *OrderItemRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, orderItem domain.OrederItem) domain.OrederItem {
	SQL := "INSERT INTO order_items(pesanan_id, produk_id,quantity,total) VALUES ($1,$2,$3,$4) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, orderItem.PesananId, orderItem.ProdukId, orderItem.Quantity, orderItem.Total).Scan(&id)
	helper.PanicIfError(err)
	orderItem.Id = id
	return orderItem
}
func (o *OrderItemRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) ([]domain.OrederItem, map[int]domain.Pesanan, map[int]domain.Pelanggan, map[int]domain.Produk) {
	SQL := "select oi.id ,oi.pesanan_id ,ps.id , ps.pelanggan_id,pl.id ,pl.name,pl.email ,pl.phone ,pl.address, ps.total_amount ,ps.status ,ps.created_at ,oi.produk_id ,p.id ,p.name ,p.price ,p.stock_quantity , oi.quantity ,oi.total from order_items oi inner join pesanans ps on oi.pesanan_id = ps.id inner join pelanggans pl on ps.pelanggan_id = pl.id inner join produks p on oi.produk_id = p.id"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var orderItems []domain.OrederItem
	pesananMap := make(map[int]domain.Pesanan)
	pelangganMap := make(map[int]domain.Pelanggan)
	produkMap := make(map[int]domain.Produk)
	for rows.Next() {
		orderItem := domain.OrederItem{}
		pesanan := domain.Pesanan{}
		pelanggan := domain.Pelanggan{}
		produk := domain.Produk{}
		err := rows.Scan(&orderItem.Id, &orderItem.PesananId, &pesanan.Id, &pesanan.PelangganId, &pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address, &pesanan.TotalAmount, &pesanan.Status, &pesanan.CreatedAt, &orderItem.ProdukId, &produk.Id, &produk.Name, &produk.Price, &produk.StockQuantity, &orderItem.Quantity, &orderItem.Total)
		helper.PanicIfError(err)
		orderItems = append(orderItems, orderItem)
		pesananMap[pesanan.Id] = pesanan
		pelangganMap[pelanggan.Id] = pelanggan
		produkMap[produk.Id] = produk
	}
	return orderItems, pesananMap, pelangganMap, produkMap
}
func (o *OrderItemRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, orderItemId int) (domain.OrederItem, domain.Pesanan, domain.Pelanggan, domain.Produk, error) {
	SQL := "select oi.id ,oi.pesanan_id ,ps.id , ps.pelanggan_id,pl.id ,pl.name,pl.email ,pl.phone ,pl.address, ps.total_amount ,ps.status ,ps.created_at ,oi.produk_id ,p.id ,p.name ,p.price ,p.stock_quantity , oi.quantity ,oi.total from order_items oi inner join pesanans ps on oi.pesanan_id = ps.id inner join pelanggans pl on ps.pelanggan_id = pl.id inner join produks p on oi.produk_id = p.id where oi.id = $1"
	rows, err := tx.QueryContext(ctx, SQL, orderItemId)
	helper.PanicIfError(err)
	defer rows.Close()

	orderItem := domain.OrederItem{}
	pesanan := domain.Pesanan{}
	pelanggan := domain.Pelanggan{}
	produk := domain.Produk{}
	if rows.Next() {
		err := rows.Scan(&orderItem.Id, &orderItem.PesananId, &pesanan.Id, &pesanan.PelangganId, &pelanggan.Id, &pelanggan.Name, &pelanggan.Email, &pelanggan.Phone, &pelanggan.Address, &pesanan.TotalAmount, &pesanan.Status, &pesanan.CreatedAt, &orderItem.ProdukId, &produk.Id, &produk.Name, &produk.Price, &produk.StockQuantity, &orderItem.Quantity, &orderItem.Total)
		helper.PanicIfError(err)
		return orderItem, pesanan, pelanggan, produk, nil
	} else {
		return orderItem, pesanan, pelanggan, produk, errors.New("order item id not found")
	}
}
func (o *OrderItemRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, orderItem domain.OrederItem) domain.OrederItem {
	SQL := "UPDATE order_items SET pesanan_id = $1, produk_id = $2,quantity = $3,total  = $4 WHERE id = $5"
	_, err := tx.ExecContext(ctx, SQL, orderItem.PesananId, orderItem.ProdukId, orderItem.Quantity, orderItem.Total, orderItem.Id)
	helper.PanicIfError(err)
	return orderItem
}
func (o *OrderItemRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, orderItem domain.OrederItem) {
	SQL := "DELETE FROM order_items WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, orderItem.Id)
	helper.PanicIfError(err)
}
