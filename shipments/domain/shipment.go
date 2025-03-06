package domain

type Shipment struct {
    ID          int    `json:"id"`
    OrderID     int    `json:"order_id"`
    TrackingID  string `json:"tracking_id"`
    Carrier     string `json:"carrier"`
    Status      string `json:"status"`
    ShipDate    string `json:"ship_date"`
    DeliveryDate string `json:"delivery_date"`
}
