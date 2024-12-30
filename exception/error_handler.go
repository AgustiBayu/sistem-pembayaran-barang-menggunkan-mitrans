package exception

import (
	"net/http"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"

	"github.com/go-playground/validator/v10"
)

func ErrorHandler(writer http.ResponseWriter, request *http.Request, err interface{}) {
	if validationError(writer, request, err) {
		return
	}
	if notFoundError(writer, request, err) {
		return
	}

	// HandleNotFound(writer, request)

	internalServerError(writer, request, err)
}

func HandleNotFound(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusNotFound)
	webResponse := web.ErrorResponse{
		Code:  http.StatusNotFound,
		Error: "end point is not valid",
	}
	helper.WriteResponseBody(writer, webResponse)
}

func validationError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(validator.ValidationErrors)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusBadRequest)

		webResponse := web.ErrorResponse{
			Code:  http.StatusBadRequest,
			Error: exception.Error(),
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}

func notFoundError(writer http.ResponseWriter, request *http.Request, err interface{}) bool {
	exception, ok := err.(NotFoundError)
	if ok {
		writer.Header().Set("Content-Type", "application/json")
		writer.WriteHeader(http.StatusNotFound)

		webResponse := web.ErrorResponse{
			Code:  http.StatusNotFound,
			Error: exception.Error,
		}
		helper.WriteResponseBody(writer, webResponse)
		return true
	} else {
		return false
	}
}
func internalServerError(writer http.ResponseWriter, request *http.Request, err interface{}) {
	writer.Header().Set("Content-Type", "application/json")
	writer.WriteHeader(http.StatusInternalServerError)

	var errorMessage string
	if err == nil {
		errorMessage = "An unexpected error occurred"
	} else {
		switch e := err.(type) {
		case string:
			errorMessage = e
		case error:
			errorMessage = e.Error()
		default:
			errorMessage = "An unexpected error occurred"
		}
	}

	webResponse := web.ErrorResponse{
		Code:  http.StatusInternalServerError,
		Error: errorMessage,
	}
	helper.WriteResponseBody(writer, webResponse)
}
