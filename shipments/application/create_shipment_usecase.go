package aplication


import ( 

	"encoding/json"
	"strconv"
	"time"

	"ModaVane/shipments/domain"
	"ModaVane/shipments/domain/ports"


)

type CreateShipmentUseCase struct {
	repo               ports.ShipmentRepository
	broker             ports.Broker
	senderNotification ports.SenderNotification
}

func NewCreateShipmentUseCase(repo ports.ShipmentRepository, broker ports.Broker, senderNotification ports.SenderNotification) *CreateShipmentUseCase {
	return &CreateShipmentUseCase{
		repo:               repo,
		broker:             broker,
		senderNotification: senderNotification,
	}
}

func (uc *CreateShipmentUseCase) Execute(shipment domain.Shipment) (int, error) {
	idShipment, err := uc.repo.CreateShipment(shipment)
	if err != nil {
		return 0, err
	}
	idShipmentStr := strconv.Itoa(idShipment)

	messageJson := map[string]interface{}{
		"shipment_id": shipment.ID,
	}

	messageJsonStr, err := json.Marshal(messageJson)
	if err != nil {
		return idShipment, err
	}

	err = uc.broker.Publish(string(messageJsonStr))
	if err != nil {
		return idShipment, err
	}
	time.Sleep(5 * time.Second)

	err = uc.senderNotification.SendNotification(map[string]interface{}{
		"event": "new-shipment",
		"data":  idShipmentStr,
	})

	if err != nil {
		return idShipment, err
	}

	return idShipment, nil
}
