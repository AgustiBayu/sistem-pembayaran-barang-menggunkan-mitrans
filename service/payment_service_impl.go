package service

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"sistem-pembayaran-barang-menggunkan-mitrans/exception"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/domain"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/repository"
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/snap"
)

type PaymentServiceImpl struct {
	PaymentRepository   repository.PaymentRepository
	PesananRepository   repository.PesananRepository
	PelangganRepository repository.PelangganRepository
	DB                  *sql.DB
	Validate            *validator.Validate
}

func NewPaymentService(
	paymentRepository repository.PaymentRepository, pesananRepository repository.PesananRepository,
	pelangganRepository repository.PelangganRepository, DB *sql.DB, validate *validator.Validate) PaymentService {
	return &PaymentServiceImpl{
		PaymentRepository:   paymentRepository,
		PesananRepository:   pesananRepository,
		PelangganRepository: pelangganRepository,
		DB:                  DB,
		Validate:            validate,
	}
}

func (service *PaymentServiceImpl) Create(ctx context.Context, request web.PaymentCreateRequest) web.PaymentMidtrans {
	// Validasi input
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)

	// Mulai transaksi database
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	// Parsing tanggal pembayaran
	PaymentDate, err := time.Parse("02-01-2006", request.PaymentDate)
	helper.PanicIfError(err)

	// Simpan data pembayaran
	payment := domain.Payment{
		PesananId:     request.PesananId,
		PaymentMethod: "mitrans",
		PaymentStatus: "pending",
		PaymentAmount: request.PaymentAmount,
		PaymentDate:   PaymentDate,
	}
	payment = service.PaymentRepository.Save(ctx, tx, payment)
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, payment.PesananId)
	helper.PanicIfError(err)
	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, pesanan.PelangganId)
	helper.PanicIfError(err)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	serverKey := os.Getenv("MIDTRANS_SERVER_KEY")
	environment := os.Getenv("MIDTRANS_ENVIRONMENT")
	if serverKey == "" {
		log.Fatal("MIDTRANS_SERVER_KEY is not set in .env")
	}
	midtrans.ServerKey = serverKey
	if environment == "production" {
		midtrans.Environment = midtrans.Production
	} else {
		midtrans.Environment = midtrans.Sandbox
	}
	snapClient := snap.Client{}
	snapClient.New(midtrans.ServerKey, midtrans.Sandbox)

	// Buat permintaan Snap API
	snapRequest := &snap.Request{
		TransactionDetails: midtrans.TransactionDetails{
			OrderID:  fmt.Sprintf("order-%d", payment.Id), // Pastikan Order ID unik
			GrossAmt: int64(request.PaymentAmount),
		},
		CustomerDetail: &midtrans.CustomerDetails{
			FName: pelanggan.Name,
			Email: pelanggan.Email,
			Phone: pelanggan.Phone,
		},
	}

	// Kirim permintaan ke Midtrans
	snapResponse, err := snapClient.CreateTransaction(snapRequest)
	if err != nil {
		fmt.Println("error snapResponse")
	}

	// Kembalikan response Midtrans
	midtransResponse := web.PaymentMidtrans{
		Token:       snapResponse.Token,
		RedirectUrl: snapResponse.RedirectURL,
	}
	return midtransResponse
}

func (service *PaymentServiceImpl) FindAll(ctx context.Context) []web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	payment, pesanan, pelanggan := service.PaymentRepository.FindAll(ctx, tx)
	return helper.ToPaymentResponses(payment, pesanan, pelanggan)
}
func (service *PaymentServiceImpl) FindById(ctx context.Context, requestId int) web.PaymentResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	payment, pesanan, pelanggan, err := service.PaymentRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	return helper.ToPaymentResponse(payment, pesanan, pelanggan)
}
func (service *PaymentServiceImpl) Update(ctx context.Context, request web.PaymentUpdateRequest) web.PaymentResponse {
	err := service.Validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	payment, _, _, err := service.PaymentRepository.FindById(ctx, tx, request.Id)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pesanan, _, err := service.PesananRepository.FindById(ctx, tx, payment.PesananId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	pelanggan, err := service.PelangganRepository.FindById(ctx, tx, pesanan.PelangganId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	PaymentDate, err := time.Parse("02-01-2006", request.PaymentDate)
	helper.PanicIfError(err)
	helper.ValidateNewDate(payment.PaymentDate, PaymentDate)

	payment.PesananId = request.PesananId
	payment.PaymentMethod = domain.StatusPaymentMethod(request.PaymentMethod)
	payment.PaymentStatus = domain.StatusPayment(request.PaymentStatus)
	payment.PaymentAmount = pesanan.TotalAmount
	payment.PaymentDate = PaymentDate

	payment = service.PaymentRepository.Update(ctx, tx, payment)
	return helper.ToPaymentResponse(payment, pesanan, pelanggan)
}
func (service *PaymentServiceImpl) Delete(ctx context.Context, requestId int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)
	defer helper.RollbackOrCommit(tx)

	payment, _, _, err := service.PaymentRepository.FindById(ctx, tx, requestId)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}
	service.PaymentRepository.Delete(ctx, tx, payment)
}
