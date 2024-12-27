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

type ProdukServiceImpl struct {
	ProdukRepository repository.ProdukRepository
	DB               *sql.DB
	Validate         *validator.Validate
}

func NewProdukService(produkRepository repository.ProdukRepository, DB *sql.DB, validate *validator.Validate) ProdukService {
	return &ProdukServiceImpl{
		ProdukRepository: produkRepository,
		DB:               DB,
		Validate:         validate,
	}
}

func (service *ProdukServiceImpl) Create(ctx context.Context, request web.ProdukCreateRequest) web.ProdukResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk := domain.Produk{
		Name:          request.Name,
		Price:         request.Price,
		StockQuantity: request.StockQuantity,
	}
	produk = service.ProdukRepository.Save(ctx, tx, produk)
	return helper.ToProdukResponse(produk)
}
func (service *ProdukServiceImpl) FindAll(ctx context.Context) []web.ProdukResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk := service.ProdukRepository.FindAll(ctx, tx)
	return helper.ToProdukResponses(produk)
}
func (service *ProdukServiceImpl) FindById(ctx context.Context, requestId int) web.ProdukResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk, err := service.ProdukRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToProdukResponse(produk)
}
func (service *ProdukServiceImpl) Update(ctx context.Context, request web.ProdukUpdateRequest) web.ProdukResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)
	produk, err := service.ProdukRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	produk.Name = request.Name
	produk.Price = request.Price
	produk.StockQuantity = request.StockQuantity

	produk = service.ProdukRepository.Update(ctx, tx, produk)
	return helper.ToProdukResponse(produk)
}
func (service *ProdukServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	produk, err := service.ProdukRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.ProdukRepository.Delete(ctx, tx, produk)
}
