package http

import (
	"api-venta/internal/adapters/domain"
	usecases "api-venta/internal/adapters/use_cases"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VentaHandler struct {
	UseCase *usecases.VentaUseCase
}

func (h *VentaHandler) GetVentaByID(c *gin.Context) {
	id := c.Param("id")

	venta, err := h.UseCase.Execute(id, domain.VentaRequest{})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Venta no encontrado"})
		return
	}
	c.JSON(http.StatusOK, venta)
}

func (h *VentaHandler) PostVentaByID(c *gin.Context) {
	id := c.Param("id")

	var req domain.VentaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Datos inválidos"})
		return
	}

	venta, err := h.UseCase.PostExecute(id, req)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Venta no encontrada"})
		return
	}
	c.JSON(http.StatusOK, venta)
}
