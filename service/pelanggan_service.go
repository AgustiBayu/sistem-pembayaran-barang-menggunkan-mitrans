package service

import (
	"context"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

type PelangganService interface {
	Create(ctx context.Context, request web.PelangganCreateRequest) web.PelangganResponse
	FindAll(ctx context.Context) []web.PelangganResponse
	FindById(ctx context.Context, requestId int) web.PelangganResponse
	Update(ctx context.Context, request web.PelangganUpdateRequest) web.PelangganResponse
	Delete(ctx context.Context, requestId int)
}
