package controller

import (
	"net/http"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type ProdukControllerImpl struct {
	ProdukService service.ProdukService
}

func NewProdukController(produkService service.ProdukService) ProdukController {
	return &ProdukControllerImpl{
		ProdukService: produkService,
	}
}

func (controller *ProdukControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkCreateRequest := web.ProdukCreateRequest{}
	helper.ReadResponseBody(request, &produkCreateRequest)

	produkResponse := controller.ProdukService.Create(request.Context(), produkCreateRequest)
	links := helper.CreateLinksForItem(produkResponse.Id, "produks")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  produkResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkResponse := controller.ProdukService.FindAll(request.Context())
	var ids []int
	for _, produk := range produkResponse {
		ids = append(ids, produk.Id)
	}
	links := helper.CreateLinksForItems(ids, "produks")
	webResponse := web.DataResponse{
		Code:  http.StatusOK,
		Data:  produkResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkId := params.ByName("produkId")
	id, err := strconv.Atoi(produkId)
	helper.PanicIfError(err)

	produkResponse := controller.ProdukService.FindById(request.Context(), id)
	links := helper.CreateLinksForItem(produkResponse.Id, "produks")
	webResponse := web.DataResponse{
		Code:  http.StatusOK,
		Data:  produkResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkUpdateRequest := web.ProdukUpdateRequest{}
	helper.ReadResponseBody(request, &produkUpdateRequest)
	produkId := params.ByName("produkId")
	id, err := strconv.Atoi(produkId)
	helper.PanicIfError(err)

	produkUpdateRequest.Id = id
	produkResponse := controller.ProdukService.Update(request.Context(), produkUpdateRequest)
	links := helper.CreateLinksForItem(produkResponse.Id, "produks")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  produkResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *ProdukControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	produkId := params.ByName("produkId")
	id, err := strconv.Atoi(produkId)
	helper.PanicIfError(err)

	controller.ProdukService.Delete(request.Context(), id)
	webResponse := web.DeleteResponse{
		Code: http.StatusNoContent,
	}
	helper.WriteResponseBody(writer, webResponse)
}
