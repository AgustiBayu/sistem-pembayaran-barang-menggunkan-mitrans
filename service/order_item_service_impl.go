package service

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/exception"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/repository"

	"github.com/go-playground/validator/v10"
)

type OrderItemServiceImpl struct {
	OrderItemRepository repository.OrderItemRepository
	PesananRepository   repository.PesananRepository
	PelangganRepository repository.PelangganRepository
	ProdukRepository    repository.ProdukRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewOrderItemService(
	orderItemRepository repository.OrderItemRepository, pesananRepository repository.PesananRepository, pelangganRepository repository.PelangganRepository,
	produkRepository repository.ProdukRepository, DB *sql.DB, validate *validator.Validate) OrderItemService {
	return &OrderItemServiceImpl{
		OrderItemRepository: orderItemRepository,
		PesananRepository:   pesananRepository,
		PelangganRepository: pelangganRepository,
		ProdukRepository:    produkRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *OrderItemServiceImpl) Create(ctx context.Context, request web.OrderItemCreateRequest) web.OrderItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	orderItem := domain.OrederItem{
		PesananId: request.PesananId,
		ProdukId:  request.ProdukId,
		Quantity:  request.Quantity,
		Total:     request.Total,
	}
	orderItem = service.OrderItemRepository.Save(ctx, tx, orderItem)
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, orderItem.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, pesanan.PelangganId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	produk, err := service.ProdukRepository.FindById(ctx, tx, orderItem.ProdukId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOrderItemResponse(orderItem, pesanan, pelanggan, produk)
}
func (service *OrderItemServiceImpl) FindAll(ctx context.Context) []web.OrderItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	orderItem, pesanan, pelanggan, produk := service.OrderItemRepository.FindAll(ctx, tx)
	return helper.ToOrderItemResponses(orderItem, pesanan, pelanggan, produk)
}
func (service *OrderItemServiceImpl) FindById(ctx context.Context, requestId int) web.OrderItemResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	orderItem, pesanan, pelanggan, produk, err := service.OrderItemRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToOrderItemResponse(orderItem, pesanan, pelanggan, produk)
}
func (service *OrderItemServiceImpl) Update(ctx context.Context, request web.OrderItemUpdateRequest) web.OrderItemResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	orderItem, _, _, _, err := service.OrderItemRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, orderItem.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, pesanan.PelangganId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	produk, err := service.ProdukRepository.FindById(ctx, tx, orderItem.ProdukId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	orderItem.PesananId = request.ProdukId
	orderItem.ProdukId = request.ProdukId
	orderItem.Quantity = request.Quantity
	orderItem.Total = request.Total

	orderItem = service.OrderItemRepository.Update(ctx, tx, orderItem)
	return helper.ToOrderItemResponse(orderItem, pesanan, pelanggan, produk)
}
func (service *OrderItemServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	orderItem, _, _, _, err := service.OrderItemRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.OrderItemRepository.Delete(ctx, tx, orderItem)
}
