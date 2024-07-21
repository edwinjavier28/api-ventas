package repository

import (
	"api-venta/internal/adapters/domain"
	"errors"
)

type VentaRepository interface {
	PostVentaByID(id string, req domain.VentaRequest) (*domain.ResponseVenta, error)
	GetVentaByID(id string) (*domain.ResponseVenta, error)
}

type InMemoryVentaRepository struct {
	venta map[string]*domain.ResponseVenta
}

func NewInMemoryVentaRepository() *InMemoryVentaRepository {
	return &InMemoryVentaRepository{
		venta: map[string]*domain.ResponseVenta{
			"1": {Idventa: "1", NombreVenta: "Perros Calientes", CostodeVenta: 300},
		},
	}
}
func (repo *InMemoryVentaRepository) GetVentaByID(id string) (*domain.ResponseVenta, error) {
	if _, ok := repo.venta[id]; ok {
		return repo.venta[id], nil
	} else {
		return nil, errors.New("error")
	}
}

func (repo *InMemoryVentaRepository) PostVentaByID(id string, req domain.VentaRequest) (*domain.ResponseVenta, error) {
	if venta, ok := repo.venta[id]; ok {
		return venta, nil
	}

	venta := &domain.ResponseVenta{
		Idventa:      id,
		NombreVenta:  req.NombreVenta,
		CostodeVenta: req.CostodeVenta,
	}

	repo.venta[id] = venta

	return venta, nil

}
