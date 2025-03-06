package adapters

import (
    "database/sql"
    "errors"

    "ModaVane/shipments/domain"

)
//k
type MySQLShipmentRepository struct {
    DB *sql.DB
}

func NewMySQLShipmentRepository(db *sql.DB) *MySQLShipmentRepository {
    return &MySQLShipmentRepository{DB: db}
}

func (repo *MySQLShipmentRepository) CreateShipment(shipment domain.Shipment) (int, error) {
    res, err := repo.DB.Exec(
        "INSERT INTO shipments (order_id, tracking_id, carrier, status, ship_date, delivery_date) VALUES (?, ?, ?, ?, ?, ?)",
        shipment.OrderID, shipment.TrackingID, shipment.Carrier, shipment.Status, shipment.ShipDate, shipment.DeliveryDate,
    )
    if err != nil {
        return 0, err
    }

    id, err := res.LastInsertId()
    if err != nil {
        return 0, err
    }

    return int(id), nil
}

func (repo *MySQLShipmentRepository) GetShipmentByID(id int) (*domain.Shipment, error) {
    var shipment domain.Shipment
    err := repo.DB.QueryRow(
        "SELECT id, order_id, tracking_id, carrier, status, ship_date, delivery_date FROM shipments WHERE id = ?",
        id,
    ).Scan(&shipment.ID, &shipment.OrderID, &shipment.TrackingID, &shipment.Carrier, &shipment.Status, &shipment.ShipDate, &shipment.DeliveryDate)

    if err != nil {
        if err == sql.ErrNoRows {
            return nil, nil
        }
        return nil, err
    }

    return &shipment, nil
}

func (repo *MySQLShipmentRepository) GetAllShipments() ([]domain.Shipment, error) {
    rows, err := repo.DB.Query("SELECT id, order_id, tracking_id, carrier, status, ship_date, delivery_date FROM shipments")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    shipments := []domain.Shipment{}
    for rows.Next() {
        var s domain.Shipment
        if err := rows.Scan(&s.ID, &s.OrderID, &s.TrackingID, &s.Carrier, &s.Status, &s.ShipDate, &s.DeliveryDate); err != nil {
            return nil, err
        }
        shipments = append(shipments, s)
    }

    if err := rows.Err(); err != nil {
        return nil, err
    }

    return shipments, nil
}

func (repo *MySQLShipmentRepository) UpdateShipment(shipment domain.Shipment) error {
    _, err := repo.DB.Exec(
        "UPDATE shipments SET order_id=?, tracking_id=?, carrier=?, status=?, ship_date=?, delivery_date=? WHERE id=?",
        shipment.OrderID, shipment.TrackingID, shipment.Carrier, shipment.Status, shipment.ShipDate, shipment.DeliveryDate, shipment.ID,
    )
    return err
}

func (repo *MySQLShipmentRepository) DeleteShipment(id int) error {
    res, err := repo.DB.Exec("DELETE FROM shipments WHERE id=?", id)
    if err != nil {
        return err
    }

    rowsAffected, err := res.RowsAffected()
    if err != nil {
        return err
    }

    if rowsAffected == 0 {
        return errors.New("no se eliminó ningún registro")
    }

    return nil
}
