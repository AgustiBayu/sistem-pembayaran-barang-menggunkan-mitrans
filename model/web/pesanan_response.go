package web

type PesananResponse struct {
	Id          int               `json:"id" validate:"required"`
	Pelanggan   PelangganResponse `json:"pelanggan" validate:"required"`
	TotalAmount int               `json:"total_amount" validate:"required"`
	Status      string            `json:"status" validate:"required"`
	CreatedAt   string            `json:"created_ad" validate:"required"`
}
