package ports

import "ModaVane/shipments/domain"

type ShipmentRepository interface {
    CreateShipment(shipment domain.Shipment) (int, error)
    GetShipmentByID(id int) (*domain.Shipment, error)
    GetAllShipments() ([]domain.Shipment, error)
    UpdateShipment(shipment domain.Shipment) error
    DeleteShipment(id int) error
}
