package service

import (
	"context"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
)

type OrderItemService interface {
	Create(ctx context.Context, request web.OrderItemCreateRequest) web.OrderItemResponse
	FindAll(ctx context.Context) []web.OrderItemResponse
	FindById(ctx context.Context, requestId int) web.OrderItemResponse
	Update(ctx context.Context, request web.OrderItemUpdateRequest) web.OrderItemResponse
	Delete(ctx context.Context, requestId int)
}
