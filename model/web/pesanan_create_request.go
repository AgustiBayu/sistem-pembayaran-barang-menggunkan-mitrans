package web

type PesananCreateRequest struct {
	PelangganId int    `json:"pelanggan_id" validate:"required"`
	TotalAmount int    `json:"total_amount" validate:"required"`
	Status      string `json:"status" validate:"required"`
	CreatedAt   string `json:"created_at" validate:"required"`
}
