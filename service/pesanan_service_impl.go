package service

import (
	"context"
	"database/sql"
	"sistem-pembayaran-barang-menggunkan-mitrans/exception"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/repository"
	"time"

	"github.com/go-playground/validator/v10"
)

type PesananServiceImpl struct {
	PesananRepository   repository.PesananRepository
	PelangganRepository repository.PelangganRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewPesananService(pesananRepository repository.PesananRepository, pelangganRepository repository.PelangganRepository,
	DB *sql.DB, validate *validator.Validate) PesananService {
	return &PesananServiceImpl{
		PesananRepository:   pesananRepository,
		PelangganRepository: pelangganRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *PesananServiceImpl) Create(ctx context.Context, request web.PesananCreateRequest) web.PesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	CreatedAt, err := time.Parse("02-01-2006", request.CreatedAt)
	helper.PanicIfError(err)
	pesanan := domain.Pesanan{
		PelangganId: request.PelangganId,
		TotalAmount: request.TotalAmount,
		Status:      domain.StatusPesanan(request.Status),
		CreatedAt:   CreatedAt,
	}
	pesanan = service.PesananRepository.Save(ctx, tx, pesanan)
	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, pesanan.PelangganId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPesananResponse(pesanan, pelanggan)
}
func (service *PesananServiceImpl) FindAll(ctx context.Context) []web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, pelanggan := service.PesananRepository.FindAll(ctx, tx)
	return helper.ToPesananResponses(pesanan, pelanggan)
}
func (service *PesananServiceImpl) FindById(ctx context.Context, requestId int) web.PesananResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, pelanggan, err := service.PesananRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPesananResponse(pesanan, pelanggan)
}
func (service *PesananServiceImpl) Update(ctx context.Context, request web.PesananUpdateRequest) web.PesananResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, pesanan.PelangganId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	CreatedAt, err := time.Parse("02-01-2006", request.CreatedAt)
	helper.PanicIfError(err)
	err = helper.ValidateNewDate(pesanan.CreatedAt, CreatedAt)
	helper.PanicIfError(err)

	pesanan.PelangganId = request.PelangganId
	pesanan.TotalAmount = request.TotalAmount
	pesanan.Status = domain.StatusPesanan(request.Status)
	pesanan.CreatedAt = CreatedAt

	return helper.ToPesananResponse(pesanan, pelanggan)
}
func (service *PesananServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PesananRepository.Delete(ctx, tx, pesanan)
}
