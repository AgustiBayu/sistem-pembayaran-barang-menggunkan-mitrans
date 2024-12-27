package helper

import (
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

func ToPaymentResponses(payments []domain.Payment, pesananMap map[int]domain.Pesanan, pelangganMap map[int]domain.Pelanggan) []web.PaymentResponse {
	var paymentResponse []web.PaymentResponse
	for _, payment := range payments {
		pesanan, exists := pesananMap[payment.PesananId]
		if !exists {
			pesanan = domain.Pesanan{}
		}
		pelanggan, exists := pelangganMap[pesanan.PelangganId]
		if !exists {
			pelanggan = domain.Pelanggan{}
		}
		paymentResponse = append(paymentResponse, ToPaymentResponse(payment, pesanan, pelanggan))
	}
	return paymentResponse
}
func ToPaymentResponse(payment domain.Payment, pesanan domain.Pesanan, pelanggan domain.Pelanggan) web.PaymentResponse {
	return web.PaymentResponse{
		Id: payment.Id,
		Pesanan: web.PesananResponse{
			Id: pesanan.Id,
			Pelanggan: web.PelangganResponse{
				Id:      pelanggan.Id,
				Name:    pelanggan.Name,
				Email:   pelanggan.Email,
				Phone:   pelanggan.Phone,
				Address: pelanggan.Address,
			},
			TotalAmount: pesanan.TotalAmount,
			Status:      string(pesanan.Status),
			CreatedAt:   FormatDate(pesanan.CreatedAt),
		},
		PaymentMethod: string(payment.PaymentMethod),
		PaymentStatus: string(payment.PaymentStatus),
		PaymentAmount: payment.PaymentAmount,
		PaymentDate:   FormatDate(payment.PaymentDate),
	}
}
func ToOrderItemResponses(orderItems []domain.OrederItem, pesananMap map[int]domain.Pesanan, pelangganMap map[int]domain.Pelanggan, produkMap map[int]domain.Produk) []web.OrderItemResponse {
	var orderItemResponse []web.OrderItemResponse
	for _, orderItem := range orderItems {
		pesanan, exists := pesananMap[orderItem.PesananId]
		if !exists {
			pesanan = domain.Pesanan{}
		}
		pelanggan, exists := pelangganMap[pesanan.PelangganId]
		if !exists {
			pelanggan = domain.Pelanggan{}
		}
		produk, exists := produkMap[orderItem.ProdukId]
		if !exists {
			produk = domain.Produk{}
		}
		orderItemResponse = append(orderItemResponse, ToOrderItemResponse(orderItem, pesanan, pelanggan, produk))
	}
	return orderItemResponse
}
func ToOrderItemResponse(orderItem domain.OrederItem, pesanan domain.Pesanan, pelanggan domain.Pelanggan, produk domain.Produk) web.OrderItemResponse {
	return web.OrderItemResponse{
		Id: orderItem.Id,
		Pesanan: web.PesananResponse{
			Id: pesanan.Id,
			Pelanggan: web.PelangganResponse{
				Id:      pelanggan.Id,
				Name:    pelanggan.Name,
				Email:   pelanggan.Email,
				Phone:   pelanggan.Phone,
				Address: pelanggan.Address,
			},
			TotalAmount: pesanan.TotalAmount,
			Status:      string(pesanan.Status),
			CreatedAt:   FormatDate(pesanan.CreatedAt),
		},
		Produk: web.ProdukResponse{
			Id:            produk.Id,
			Name:          produk.Name,
			Price:         produk.Price,
			StockQuantity: produk.StockQuantity,
		},
		Quantity: orderItem.Quantity,
		Total:    orderItem.Total,
	}
}

func ToPesananResponses(pesanans []domain.Pesanan, pelangganMap map[int]domain.Pelanggan) []web.PesananResponse {
	var pesananResponse []web.PesananResponse
	for _, pesanan := range pesanans {
		pelanggan, exists := pelangganMap[pesanan.PelangganId]
		if !exists {
			pelanggan = domain.Pelanggan{}
		}
		pesananResponse = append(pesananResponse, ToPesananResponse(pesanan, pelanggan))
	}
	return pesananResponse
}

func ToPesananResponse(pesanan domain.Pesanan, pelanggan domain.Pelanggan) web.PesananResponse {
	return web.PesananResponse{
		Id: pesanan.Id,
		Pelanggan: web.PelangganResponse{
			Id:      pelanggan.Id,
			Name:    pelanggan.Name,
			Email:   pelanggan.Email,
			Phone:   pelanggan.Phone,
			Address: pelanggan.Address,
		},
		TotalAmount: pesanan.TotalAmount,
		Status:      string(pesanan.Status),
		CreatedAt:   FormatDate(pesanan.CreatedAt),
	}
}

func ToProdukResponses(produks []domain.Produk) []web.ProdukResponse {
	var produkResponse []web.ProdukResponse
	for _, produk := range produks {
		produkResponse = append(produkResponse, ToProdukResponse(produk))
	}
	return produkResponse
}

func ToProdukResponse(produk domain.Produk) web.ProdukResponse {
	return web.ProdukResponse{
		Id:            produk.Id,
		Name:          produk.Name,
		Price:         produk.Price,
		StockQuantity: produk.StockQuantity,
	}
}

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
