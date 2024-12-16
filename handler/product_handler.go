package handler

import (
	"go-wire-example/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	Service *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"products": []any{}})
}
