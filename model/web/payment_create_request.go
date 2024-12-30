package web

type PaymentCreateRequest struct {
	PesananId     int    `json:"pesanan_id" validate:"required"`
	PaymentAmount int    `json:"payment_amount" validate:"required"`
	PaymentDate   string `json:"payment_date" validate:"required"`
}
