package controller

import (
	"net/http"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PaymentControllerImpl struct {
	PaymentService service.PaymentService
}

func NewPaymentController(paymentService service.PaymentService) PaymentController {
	return &PaymentControllerImpl{
		PaymentService: paymentService,
	}
}

func (controller *PaymentControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentCreateRequest := web.PaymentCreateRequest{}
	helper.ReadResponseBody(request, &paymentCreateRequest)

	paymentResponse := controller.PaymentService.Create(request.Context(), paymentCreateRequest)
	webResponse := web.MidtransResponse{
		Code: http.StatusOK,
		Data: paymentResponse,
	}
	helper.WriteResponseBody(writer, webResponse)
}

func (controller *PaymentControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentResponse := controller.PaymentService.FindAll(request.Context())
	var ids []int
	for _, payment := range paymentResponse {
		ids = append(ids, payment.Id)
	}
	links := helper.CreateLinksForItems(ids, "payments")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  paymentResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PaymentControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentId := params.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	paymentResponse := controller.PaymentService.FindById(request.Context(), id)
	links := helper.CreateLinksForItem(paymentResponse.Id, "payments")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  paymentResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PaymentControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentUpdateRequest := web.PaymentUpdateRequest{}
	helper.ReadResponseBody(request, &paymentUpdateRequest)
	paymentId := params.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	paymentUpdateRequest.Id = id
	paymentResponse := controller.PaymentService.Update(request.Context(), paymentUpdateRequest)
	links := helper.CreateLinksForItem(paymentResponse.Id, "payments")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  paymentResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PaymentControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	paymentId := params.ByName("paymentId")
	id, err := strconv.Atoi(paymentId)
	helper.PanicIfError(err)

	controller.PaymentService.Delete(request.Context(), id)
	webResponse := web.DeleteResponse{
		Code: http.StatusCreated,
	}
	helper.WriteResponseBody(writer, webResponse)
}
