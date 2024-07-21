package main

import (
	"api-venta/internal/adapters/http"
	usecases "api-venta/internal/adapters/use_cases"

	"api-venta/internal/adapters/repository"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Crear el repositorio
	productoRepo := repository.NewInMemoryVentaRepository()

	// Crear el caso de uso
	getProductoUseCase := &usecases.VentaUseCase{
		Repo: productoRepo,
	}

	// Crear el handler
	VentasHandler := &http.VentaHandler{
		UseCase: getProductoUseCase,
	}

	// Definir la ruta

	r.POST("/Ventas/:id", VentasHandler.PostVentaByID)
	r.GET("/Ventas/:id", VentasHandler.GetVentaByID)
	// Iniciar el servidor
	r.Run(":8089")
}
