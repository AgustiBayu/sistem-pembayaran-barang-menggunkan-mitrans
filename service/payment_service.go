package service

import (
	"context"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

type PaymentService interface {
	Create(ctx context.Context, request web.PaymentCreateRequest) web.PaymentResponse
	FindAll(ctx context.Context) []web.PaymentResponse
	FindById(ctx context.Context, requestId int) web.PaymentResponse
	Update(ctx context.Context, request web.PaymentUpdateRequest) web.PaymentResponse
	Delete(ctx context.Context, requestId int)
}
