package controller

import (
	"net/http"
	"sistem-pembayaran-barang-menggunkan-mitrans/helper"
	"sistem-pembayaran-barang-menggunkan-mitrans/model/web"
	"sistem-pembayaran-barang-menggunkan-mitrans/service"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type PelangganControllerImpl struct {
	PelangganService service.PelangganService
}

func NewPelangganController(pelangganService service.PelangganService) PelangganController {
	return &PelangganControllerImpl{
		PelangganService: pelangganService,
	}
}

func (controller *PelangganControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pelangganCreateRequest := web.PelangganCreateRequest{}
	helper.ReadResponseBody(request, &pelangganCreateRequest)

	pelangganResponse := controller.PelangganService.Create(request.Context(), pelangganCreateRequest)
	links := helper.CreateLinksForItem(pelangganResponse.Id, "pelanggans")

	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  pelangganResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PelangganControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pelangganResponse := controller.PelangganService.FindAll(request.Context())
	var ids []int
	for _, pelanggan := range pelangganResponse {
		ids = append(ids, pelanggan.Id)
	}
	links := helper.CreateLinksForItems(ids, "pelanggans")
	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  pelangganResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PelangganControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pelangganId := params.ByName("pelangganId")
	id, err := strconv.Atoi(pelangganId)
	helper.PanicIfError(err)
	pelangganResponse := controller.PelangganService.FindById(request.Context(), id)
	links := helper.CreateLinksForItem(pelangganResponse.Id, "pelanggans")

	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  pelangganResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PelangganControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pelangganUpdateRequest := web.PelangganUpdateRequest{}
	helper.ReadResponseBody(request, &pelangganUpdateRequest)
	pelangganId := params.ByName("pelangganId")
	id, err := strconv.Atoi(pelangganId)
	helper.PanicIfError(err)

	pelangganUpdateRequest.Id = id
	pelangganResponse := controller.PelangganService.Update(request.Context(), pelangganUpdateRequest)
	links := helper.CreateLinksForItem(pelangganResponse.Id, "pelanggans")

	webResponse := web.DataResponse{
		Code:  http.StatusCreated,
		Data:  pelangganResponse,
		Links: links,
	}
	helper.WriteResponseBody(writer, webResponse)
}
func (controller *PelangganControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	pelangganId := params.ByName("pelangganId")
	id, err := strconv.Atoi(pelangganId)
	helper.PanicIfError(err)

	controller.PelangganService.Delete(request.Context(), id)
	webResponse := web.DeleteResponse{
		Code: http.StatusCreated,
	}
	helper.WriteResponseBody(writer, webResponse)
}
