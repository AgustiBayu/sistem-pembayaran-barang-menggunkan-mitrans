package web

type OrderItemUpdateRequest struct {
	Id        int `json:"id" validate:"required"`
	PesananId int `json:"pesanan_id" validate:"required"`
	ProdukId  int `json:"produk_id" validate:"required"`
	Quantity  int `json:"quantity" validate:"required"`
	Total     int `json:"total" validate:"required"`
}
