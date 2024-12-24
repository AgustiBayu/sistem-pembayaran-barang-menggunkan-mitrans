package app

import (
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"time"
)

func NewDB() *sql.DB {
	db, err := sql.Open("postgres", "postgres://postgres:terserah123@localhost:5432/db_sistempembayaran?sslmode=disable")
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
