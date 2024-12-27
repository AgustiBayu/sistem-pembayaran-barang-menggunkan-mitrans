package controller

import (
	"net/http"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type OrderItemControllerImpl struct {
	OrderItemService service.OrderItemService
}

func NewOrderItemController(orderItemService service.OrderItemService) OrderItemController {
	return &OrderItemControllerImpl{
		OrderItemService: orderItemService,
	}
}

func (controller *OrderItemControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderItemCreateRequest := web.OrderItemCreateRequest{}
	helper.ReadResponseBody(request, &orderItemCreateRequest)

	orderItemResponse := controller.OrderItemService.Create(request.Context(), orderItemCreateRequest)
	links := helper.CreateLinksForItem(orderItemResponse.Id, "orderItems")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  orderItemResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *OrderItemControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderItemResponse := controller.OrderItemService.FindAll(request.Context())
	var ids []int
	for _, orderItem := range orderItemResponse {
		ids = append(ids, orderItem.Id)
	}
	links := helper.CreateLinksForItems(ids, "orderItems")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  orderItemResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *OrderItemControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderItemId := params.ByName("orderItemId")
	id, err := strconv.Atoi(orderItemId)
	helper.PanicIfError(err)

	orderItemResponse := controller.OrderItemService.FindById(request.Context(), id)
	links := helper.CreateLinksForItem(orderItemResponse.Id, "orderItems")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  orderItemResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *OrderItemControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderItemUpdateRequest := web.OrderItemUpdateRequest{}
	helper.ReadResponseBody(request, &orderItemUpdateRequest)
	orderItemId := params.ByName("orderItemId")
	id, err := strconv.Atoi(orderItemId)
	helper.PanicIfError(err)

	orderItemUpdateRequest.Id = id
	orderItemResponse := controller.OrderItemService.Update(request.Context(), orderItemUpdateRequest)
	links := helper.CreateLinksForItem(orderItemResponse.Id, "orderItems")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  orderItemResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *OrderItemControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	orderItemId := params.ByName("orderItemId")
	id, err := strconv.Atoi(orderItemId)
	helper.PanicIfError(err)

	controller.OrderItemService.Delete(request.Context(), id)
	webResponse := web.DeleteResponse{
		Code: http.StatusCreated,
	}
	helper.WriteResponseBody(writer, webResponse)
}
