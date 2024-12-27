package web

type OrderItemResponse struct {
	Id       int             `json:"id" validate:"required"`
	Pesanan  PesananResponse `json:"pesanan" validate:"required"`
	Produk   ProdukResponse  `json:"produk" validate:"required"`
	Quantity int             `json:"quantity" validate:"required"`
	Total    int             `json:"total" validate:"required"`
}
