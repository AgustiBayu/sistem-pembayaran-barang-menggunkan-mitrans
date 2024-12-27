package web

type ProdukCreateRequest struct {
	Name          string `json:"name" validate:"required"`
	Price         int    `json:"price" validate:"required"`
	StockQuantity int    `json:"stock_quantity" validate:"required"`
}
