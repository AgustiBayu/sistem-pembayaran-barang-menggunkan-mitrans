package web

type PelangganUpdateRequest struct {
	Id      int    `json:"id" validate:"required"`
	Name    string `json:"name" validate:"required"`
	Email   string `json:"email" validate:"required"`
	Phone   string `json:"phone" validate:"required"`
	Address string `json:"address" validate:"required"`
}
