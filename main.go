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

	router := httprouter.New()
	router.POST("/api/pelanggans", pelangganController.Create)
	router.GET("/api/pelanggans", pelangganController.FindAll)
	router.GET("/api/pelanggans/:pelangganId", pelangganController.FindById)
	router.PUT("/api/pelanggans/:pelangganId", pelangganController.Update)
	router.DELETE("/api/pelanggans/:pelangganId", pelangganController.Delete)

	router.NotFound = http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		exception.HandleNotFound(writer, request)
	})
	router.PanicHandler = exception.ErrorHandler
	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
