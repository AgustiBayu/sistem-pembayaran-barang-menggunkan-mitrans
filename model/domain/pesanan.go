package domain

import "time"

type StatusPesanan string

const (
	StatusPending   StatusPesanan = "pending"
	StatusCompleted StatusPesanan = "completed"
	StatusCancelled StatusPesanan = "cancelled"
)

type Pesanan struct {
	Id          int
	PelangganId int
	TotalAmount int
	Status      StatusPesanan
	CreatedAt   time.Time
}
