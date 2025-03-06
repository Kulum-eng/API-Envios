package aplication

import (
    "ModaVane/shipments/domain"
    "ModaVane/shipments/domain/ports"
)

type CreateShipmentUseCase struct {
    repo ports.ShipmentRepository
}

func NewCreateShipmentUseCase(repo ports.ShipmentRepository) *CreateShipmentUseCase {
    return &CreateShipmentUseCase{repo: repo}
}

func (uc *CreateShipmentUseCase) Execute(shipment domain.Shipment) (int, error) {
    return uc.repo.CreateShipment(shipment)
}
