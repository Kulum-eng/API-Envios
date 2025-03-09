package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	s_application "ModaVane/shipments/application"
	core "ModaVane/shipments/core"
	s_adapters "ModaVane/shipments/infraestructure/adapters"
	s_controllers "ModaVane/shipments/infraestructure/http/controllers"
	s_routes "ModaVane/shipments/infraestructure/http/routes"

)

func CORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("CORS")
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	myGin := gin.New()
	myGin.RedirectTrailingSlash = false

	myGin.Use(CORS())

	db, err := core.InitDB()
	if err != nil {
		log.Println("Error al conectar a la base de datos:", err)
		return
	}

	rabbitBroker := s_adapters.NewRabbitMQBroker("ec2-3-83-91-51.compute-1.amazonaws.com", 5672, "ale", "ale123")

	err = rabbitBroker.Connect()
	if err != nil {
		log.Println("Error al conectar a RabbitMQ:", err)
		return
	}

	err = rabbitBroker.InitChannel("envios")
	if err != nil {
		log.Println("Error al inicializar el canal de RabbitMQ:", err)
		return
	}

	shipmentRepository := s_adapters.NewMySQLShipmentRepository(db)
	httpSenderNotification := s_adapters.NewHTTPSenderNotification("localhost", 3000)

	createShipmentUseCase := s_application.NewCreateShipmentUseCase(shipmentRepository, rabbitBroker, httpSenderNotification)
	getShipmentUseCase := s_application.NewGetShipmentUseCase(shipmentRepository)
	updateShipmentUseCase := s_application.NewUpdateShipmentUseCase(shipmentRepository)
	deleteShipmentUseCase := s_application.NewDeleteShipmentUseCase(shipmentRepository)

	createShipmentController := s_controllers.NewShipmentController(createShipmentUseCase, getShipmentUseCase, updateShipmentUseCase, deleteShipmentUseCase)
	s_routes.SetupShipmentRoutes(myGin, createShipmentController)

	if err := myGin.Run(":8086"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
