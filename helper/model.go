package helper

import (
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

func ToPelangganResponses(pelanggans []domain.Pelanggan) []web.PelangganResponse {
	var pelangganResponse []web.PelangganResponse
	for _, pelanggan := range pelanggans {
		pelangganResponse = append(pelangganResponse, ToPelangganResponse(pelanggan))
	}
	return pelangganResponse
}

func ToPelangganResponse(pelanggan domain.Pelanggan) web.PelangganResponse {
	return web.PelangganResponse{
		Id:      pelanggan.Id,
		Name:    pelanggan.Name,
		Email:   pelanggan.Email,
		Phone:   pelanggan.Phone,
		Address: pelanggan.Address,
	}
}
