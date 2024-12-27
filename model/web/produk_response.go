package web

type ProdukResponse struct {
	Id            int    `json:"id" validate:"required"`
	Name          string `json:"name" validate:"required"`
	Price         int    `json:"price" validate:"required"`
	StockQuantity int    `json:"stock_quantity" validate:"required"`
}
