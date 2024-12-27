package repository

import (
	"context"
	"database/sql"
	"errors"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
)

type ProdukRepositoryImpl struct{}

func NewProdukRepository() ProdukRepository {
	return &ProdukRepositoryImpl{}
}

func (p *ProdukRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, produk domain.Produk) domain.Produk {
	SQL := "INSERT INTO produks(name, price, stock_quantity) VALUES($1,$2,$3) RETURNING id"
	var id int
	err := tx.QueryRowContext(ctx, SQL, produk.Name, produk.Price, produk.StockQuantity).Scan(&id)
	helper.PanicIfError(err)
	produk.Id = id
	return produk
}
func (p *ProdukRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Produk {
	SQL := "SELECT id, name, price, stock_quantity FROM produks"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var produks []domain.Produk
	for rows.Next() {
		produk := domain.Produk{}
		err := rows.Scan(&produk.Id, &produk.Name, &produk.Price, &produk.StockQuantity)
		helper.PanicIfError(err)
		produks = append(produks, produk)
	}
	return produks
}
func (p *ProdukRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, produkId int) (domain.Produk, error) {
	SQL := "SELECT id, name, price, stock_quantity FROM produks WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, produkId)
	helper.PanicIfError(err)
	defer rows.Close()

	produk := domain.Produk{}
	if rows.Next() {
		err := rows.Scan(&produk.Id, &produk.Name, &produk.Price, &produk.StockQuantity)
		helper.PanicIfError(err)
		return produk, nil
	} else {
		return produk, errors.New("produk id not found")
	}
}
func (p *ProdukRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, produk domain.Produk) domain.Produk {
	SQL := "UPDATE produks SET name = $1, price = $2, stock_quantity = $3 WHERE id = $4"
	_, err := tx.ExecContext(ctx, SQL, produk.Name, produk.Price, produk.StockQuantity, produk.Id)
	helper.PanicIfError(err)
	return produk
}
func (p *ProdukRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, produk domain.Produk) {
	SQL := "DELETE FROM produks WHERE id = $1"
	_, err := tx.ExecContext(ctx, SQL, produk.Id)
	helper.PanicIfError(err)
}
