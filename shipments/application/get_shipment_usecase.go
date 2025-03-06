package aplication

import (
    "ModaVane/shipments/domain"
    "ModaVane/shipments/domain/ports"
)

type GetShipmentUseCase struct {
    repo ports.ShipmentRepository
}

func NewGetShipmentUseCase(repo ports.ShipmentRepository) *GetShipmentUseCase {
    return &GetShipmentUseCase{repo: repo}
}

func (uc *GetShipmentUseCase) ExecuteByID(id int) (*domain.Shipment, error) {
    return uc.repo.GetShipmentByID(id)
}

func (uc *GetShipmentUseCase) ExecuteAll() ([]domain.Shipment, error) {
    return uc.repo.GetAllShipments()
}
