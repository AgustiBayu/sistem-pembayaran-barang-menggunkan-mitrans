package service

import (
	"context"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

type PesananService interface {
	Create(ctx context.Context, request web.PesananCreateRequest) web.PesananResponse
	FindAll(ctx context.Context) []web.PesananResponse
	FindById(ctx context.Context, requestId int) web.PesananResponse
	Update(ctx context.Context, request web.PesananUpdateRequest) web.PesananResponse
	Delete(ctx context.Context, requestId int)
}
