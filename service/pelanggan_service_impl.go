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

type PelangganServiceImpl struct {
	PelangganRepository repository.PelangganRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewPelangganService(pelangganRepository repository.PelangganRepository, DB *sql.DB, validate *validator.Validate) PelangganService {
	return &PelangganServiceImpl{
		PelangganRepository: pelangganRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *PelangganServiceImpl) Create(ctx context.Context, request web.PelangganCreateRequest) web.PelangganResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pelanggan := domain.Pelanggan{
		Name:    request.Name,
		Email:   request.Email,
		Phone:   request.Phone,
		Address: request.Address,
	}
	pelanggan = service.PelangganRepository.Save(ctx, tx, pelanggan)
	return helper.ToPelangganResponse(pelanggan)
}
func (service *PelangganServiceImpl) FindAll(ctx context.Context) []web.PelangganResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pelanggan := service.PelangganRepository.FindAll(ctx, tx)
	return helper.ToPelangganResponses(pelanggan)
}
func (service *PelangganServiceImpl) FindById(ctx context.Context, requestId int) web.PelangganResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPelangganResponse(pelanggan)
}
func (service *PelangganServiceImpl) Update(ctx context.Context, request web.PelangganUpdateRequest) web.PelangganResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pelanggan.Name = request.Name
	pelanggan.Email = request.Email
	pelanggan.Phone = request.Phone
	pelanggan.Address = request.Address

	pelanggan = service.PelangganRepository.Update(ctx, tx, pelanggan)
	return helper.ToPelangganResponse(pelanggan)
}
func (service *PelangganServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PelangganRepository.Delete(ctx, tx, pelanggan)
}
