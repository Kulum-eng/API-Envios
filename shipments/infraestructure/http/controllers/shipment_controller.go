package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	aplication "ModaVane/shipments/application"
	"ModaVane/shipments/domain"
	"ModaVane/shipments/infraestructure/http/responses"
)

type ShipmentController struct {
	createShipmentUseCase *aplication.CreateShipmentUseCase
	getShipmentUseCase    *aplication.GetShipmentUseCase
	updateShipmentUseCase *aplication.UpdateShipmentUseCase
	deleteShipmentUseCase *aplication.DeleteShipmentUseCase
}

func NewShipmentController(createUC *aplication.CreateShipmentUseCase, getUC *aplication.GetShipmentUseCase, updateUC *aplication.UpdateShipmentUseCase, deleteUC *aplication.DeleteShipmentUseCase) *ShipmentController {
	return &ShipmentController{
		createShipmentUseCase: createUC,
		getShipmentUseCase:    getUC,
		updateShipmentUseCase: updateUC,
		deleteShipmentUseCase: deleteUC,
	}
}

func (ctrl *ShipmentController) Create(ctx *gin.Context) {
	var shipment domain.Shipment
	if err := ctx.ShouldBindJSON(&shipment); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("los datos son inválidos", err.Error()))
		return
	}

	idShipment, err := ctrl.createShipmentUseCase.Execute(shipment)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al crear envío", err.Error()))
		return
	}

	shipment.ID = idShipment
	ctx.JSON(http.StatusCreated, responses.SuccessResponse("Envío creado exitosamente", shipment))
}

func (ctrl *ShipmentController) GetAll(ctx *gin.Context) {
	shipments, err := ctrl.getShipmentUseCase.ExecuteAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener envíos", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Envíos obtenidos exitosamente", shipments))
}

func (ctrl *ShipmentController) GetByID(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	shipment, err := ctrl.getShipmentUseCase.ExecuteByID(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al obtener envío", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Envío obtenido exitosamente", shipment))
}

func (ctrl *ShipmentController) Update(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	var shipment domain.Shipment
	if err := ctx.ShouldBindJSON(&shipment); err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("Datos inválidos", err.Error()))
		return
	}

	shipment.ID = id
	if err := ctrl.updateShipmentUseCase.Execute(shipment); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al actualizar envío", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Envío actualizado exitosamente", shipment))
}

func (ctrl *ShipmentController) Delete(ctx *gin.Context) {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, responses.ErrorResponse("ID inválido", err.Error()))
		return
	}

	if err := ctrl.deleteShipmentUseCase.Execute(id); err != nil {
		ctx.JSON(http.StatusInternalServerError, responses.ErrorResponse("Error al eliminar envío", err.Error()))
		return
	}

	ctx.JSON(http.StatusOK, responses.SuccessResponse("Envío eliminado exitosamente", nil))
}
