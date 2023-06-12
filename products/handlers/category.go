package handlers

import (
	"github.com/iamuditg/services"
	"net/http"
)

type CategoryHandler struct {
	services services.CategoryService
}

func NewCategoryHandler(service services.CategoryService) *CategoryHandler {
	return &CategoryHandler{services: service}
}

func (c *CategoryHandler) CreateCategoryHandler(writer http.ResponseWriter, request *http.Request) {

}
