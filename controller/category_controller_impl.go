package controller

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/qulDev/golang-restful-api/service"
)

type CategoryControllerImpl struct {
	CategoryService service.CategoryService
}

func (controller *CategoryControllerImpl) Create(w http.ResponseWriter, r http.Request, params httprouter.Params) {
	//TODO implement me

	panic("implement me")
}

func (controller *CategoryControllerImpl) Update(w http.ResponseWriter, r http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryControllerImpl) Delete(w http.ResponseWriter, r http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryControllerImpl) FindById(w http.ResponseWriter, r http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *CategoryControllerImpl) FindAll(w http.ResponseWriter, r http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
