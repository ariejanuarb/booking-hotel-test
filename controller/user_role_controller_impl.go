package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"programmerzamannow/belajar-golang-restful-api/helper"
	"programmerzamannow/belajar-golang-restful-api/model/web"
	"programmerzamannow/belajar-golang-restful-api/service"
	"strconv"
)

type UserRoleControllerImpl struct {
	UserRoleService service.UserRoleService
}

func NewUserRoleController(userRoleService service.UserRoleService) UserRoleController {
	return &UserRoleControllerImpl{
		UserRoleService: userRoleService,
	}
}

func (controller *UserRoleControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRoleCreateRequest := web.UserRoleCreateRequest{}
	helper.ReadFromRequestBody(request, &userRoleCreateRequest)

	userRoleResponse := controller.UserRoleService.Create(request.Context(), userRoleCreateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userRoleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserRoleControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRoleUpdateRequest := web.UserRoleUpdateRequest{}
	helper.ReadFromRequestBody(request, &userRoleUpdateRequest)

	userRoleId := params.ByName("userRoleId")
	id, err := strconv.Atoi(userRoleId)
	helper.PanicIfError(err)

	userRoleUpdateRequest.Id = id

	userRoleResponse := controller.UserRoleService.Update(request.Context(), userRoleUpdateRequest)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userRoleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserRoleControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRoleId := params.ByName("userRoleId")
	id, err := strconv.Atoi(userRoleId)
	helper.PanicIfError(err)

	controller.UserRoleService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserRoleControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRoleId := params.ByName("userRoleId")
	id, err := strconv.Atoi(userRoleId)
	helper.PanicIfError(err)

	userRoleResponse := controller.UserRoleService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userRoleResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *UserRoleControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	userRoleResponses := controller.UserRoleService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   userRoleResponses,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
