package usecases

import (
	"api-venta/internal/adapters/domain"
	"api-venta/internal/adapters/repository"
)

type VentaUseCase struct {
	Repo repository.VentaRepository
}

func (uc *VentaUseCase) Execute(id string, req domain.VentaRequest) (*domain.ResponseVenta, error) {
	return uc.Repo.GetVentaByID(id)
}

func (uc *VentaUseCase) PostExecute(id string, req domain.VentaRequest) (*domain.ResponseVenta, error) {
	return uc.Repo.PostVentaByID(id, req)
}
