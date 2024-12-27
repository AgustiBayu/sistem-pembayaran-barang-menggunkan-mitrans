package web

type PaymentResponse struct {
	Id            int             `json:"id" validate:"required"`
	Pesanan       PesananResponse `json:"pesanan" validate:"required"`
	PaymentMethod string          `json:"payment_method" validate:"required"`
	PaymentStatus string          `json:"payment_status" validate:"required"`
	PaymentAmount int             `json:"payment_amount" validate:"required"`
	PaymentDate   string          `json:"payment_date" validate:"required"`
}
