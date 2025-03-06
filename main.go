package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	s_application "ModaVane/shipments/application"
	s_adapters "ModaVane/shipments/infraestructure/adapters"
	s_controllers "ModaVane/shipments/infraestructure/http/controllers"
	s_routes "ModaVane/shipments/infraestructure/http/routes"
	s_core "ModaVane/shipments/core"

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
	// Deshabilitar la redirección automática de barras diagonales
	gin.SetMode(gin.ReleaseMode)
	myGin := gin.New()
	myGin.RedirectTrailingSlash = false

	myGin.Use(CORS())

	db, err := core.InitDB()
	if err != nil {
		log.Println(err)
		return
	}

	// Configuración de envíos
	shipmentRepository := s_adapters.NewMySQLShipmentRepository(db)
	createShipmentUseCase := s_application.NewCreateShipmentUseCase(shipmentRepository)
	getShipmentUseCase := s_application.NewGetShipmentUseCase(shipmentRepository)
	updateShipmentUseCase := s_application.NewUpdateShipmentUseCase(shipmentRepository)
	deleteShipmentUseCase := s_application.NewDeleteShipmentUseCase(shipmentRepository)

	createShipmentController := s_controllers.NewShipmentController(createShipmentUseCase, getShipmentUseCase, updateShipmentUseCase, deleteShipmentUseCase)
	s_routes.SetupShipmentRoutes(myGin, createShipmentController)

	myGin.Run(":8080")
}
