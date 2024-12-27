package main

import (
	"net/http"
	"sistem-pembayaran-barang-menggunkan-mitrans/app"
	"sistem-pembayaran-barang-menggunkan-mitrans/controller"
	"sistem-pembayaran-barang-menggunkan-mitrans/exception"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/repository"
	"sistem-pembayaran-barang-menggunkan-mitrans/service"

	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	_ "github.com/lib/pq"
)

func main() {
	validate := validator.New()
	db := app.NewDB()

	pelangganRepository := repository.NewPelangganRepository()
	pelangganService := service.NewPelangganService(pelangganRepository, db, validate)
	pelangganController := controller.NewPelangganController(pelangganService)

	pesananRepository := repository.NewPesananRepository()
	pesananService := service.NewPesananService(pesananRepository, pelangganRepository, db, validate)
	pesananController := controller.NewPesananController(pesananService)

	produkRepository := repository.NewProdukRepository()
	produkService := service.NewProdukService(produkRepository, db, validate)
	produkController := controller.NewProdukController(produkService)

	paymentRepository := repository.NewPaymentRepository()
	paymentService := service.NewPaymentService(paymentRepository, pesananRepository, pelangganRepository, db, validate)
	paymentController := controller.NewPaymentController(paymentService)

	orderItemRepository := repository.NewOrderItemRepository()
	orderItemService := service.NewOrderItemService(orderItemRepository, pesananRepository, pelangganRepository, produkRepository, db, validate)
	orderItemController := controller.NewOrderItemController(orderItemService)

	router := httprouter.New()
	router.POST("/api/pelanggans", pelangganController.Create)
	router.GET("/api/pelanggans", pelangganController.FindAll)
	router.GET("/api/pelanggans/:pelangganId", pelangganController.FindById)
	router.PUT("/api/pelanggans/:pelangganId", pelangganController.Update)
	router.DELETE("/api/pelanggans/:pelangganId", pelangganController.Delete)

	router.POST("/api/pesanans", pesananController.Create)
	router.GET("/api/pesanans", pesananController.FindAll)
	router.GET("/api/pesanans/:pesananId", pesananController.FindById)
	router.PUT("/api/pesanans/:pesananId", pesananController.Update)
	router.DELETE("/api/pesanans/:pesananId", pesananController.Delete)

	router.POST("/api/produks", produkController.Create)
	router.GET("/api/produks", produkController.FindAll)
	router.GET("/api/produks/:produkId", produkController.FindById)
	router.PUT("/api/produks/:produkId", produkController.Update)
	router.DELETE("/api/produks/:produkId", produkController.Delete)

	router.POST("/api/payments", paymentController.Create)
	router.GET("/api/payments", paymentController.FindAll)
	router.GET("/api/payments/:paymentId", paymentController.FindById)
	router.PUT("/api/payments/:paymentId", paymentController.Update)
	router.DELETE("/api/payments/:paymentId", paymentController.Delete)

	router.POST("/api/orderItems", orderItemController.Create)
	router.GET("/api/orderItems", orderItemController.FindAll)
	router.GET("/api/orderItems/:orderItemId", orderItemController.FindById)
	router.PUT("/api/orderItems/:orderItemId", orderItemController.Update)
	router.DELETE("/api/orderItems/:orderItemId", orderItemController.Delete)

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		exception.HandleNotFound(writer, request)
	})
	router.PanicHandler = func(writer http.ResponseWriter, request *http.Request, err interface{}) {
		exception.ErrorHandler(writer, request, err)
	}

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
