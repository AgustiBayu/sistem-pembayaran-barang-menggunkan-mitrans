package service

import (
	"context"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

type ProdukService interface {
	Create(ctx context.Context, request web.ProdukCreateRequest) web.ProdukResponse
	FindAll(ctx context.Context) []web.ProdukResponse
	FindById(ctx context.Context, requestId int) web.ProdukResponse
	Update(ctx context.Context, request web.ProdukUpdateRequest) web.ProdukResponse
	Delete(ctx context.Context, requestId int)
}
