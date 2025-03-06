package aplication

import (
    "ModaVane/shipments/domain"
    "ModaVane/shipments/domain/ports"
)

type UpdateShipmentUseCase struct {
    repo ports.ShipmentRepository
}

func NewUpdateShipmentUseCase(repo ports.ShipmentRepository) *UpdateShipmentUseCase {
    return &UpdateShipmentUseCase{repo: repo}
}

func (uc *UpdateShipmentUseCase) Execute(shipment domain.Shipment) error {
    return uc.repo.UpdateShipment(shipment)
}
