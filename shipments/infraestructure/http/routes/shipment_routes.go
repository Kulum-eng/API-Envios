package routes

import (
    "github.com/gin-gonic/gin"
    "ModaVane/shipments/infraestructure/http/controllers"
)

func SetupShipmentRoutes(router *gin.Engine, controller *controllers.ShipmentController) {
    shipmentRoutes := router.Group("/shipments")
    {
        shipmentRoutes.POST("/", controller.Create)
        shipmentRoutes.GET("/", controller.GetAll)
        shipmentRoutes.GET("/:id", controller.GetByID)
        shipmentRoutes.PUT("/:id", controller.Update)
        shipmentRoutes.DELETE("/:id", controller.Delete)
    }
}
