package aplication

import (
    "ModaVane/shipments/domain/ports"
)

type DeleteShipmentUseCase struct {
    repo ports.ShipmentRepository
}

func NewDeleteShipmentUseCase(repo ports.ShipmentRepository) *DeleteShipmentUseCase {
    return &DeleteShipmentUseCase{repo: repo}
}

func (uc *DeleteShipmentUseCase) Execute(id int) error {
    return uc.repo.DeleteShipment(id)
}
