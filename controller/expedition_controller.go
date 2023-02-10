package controller

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/web"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/dimassfeb-09/restapi-ecommerce.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ExpeditionController interface {
	AddExpedition(c *gin.Context)
	UpdateExpedition(c *gin.Context)
	DeleteExpedition(c *gin.Context)
	FindAllExpedition(c *gin.Context)
	FindExpeditionByID(c *gin.Context)
}

type ExpeditionControllerImpl struct {
	ExpeditionService services.ExpeditionService
}

func NewExpeditionControllerImpl(expeditionService services.ExpeditionService) ExpeditionController {
	return &ExpeditionControllerImpl{ExpeditionService: expeditionService}
}

func (e *ExpeditionControllerImpl) AddExpedition(c *gin.Context) {
	var createRequest *request.CreateExpeditionRequest
	if err := c.ShouldBindJSON(&createRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if msgErr := helpers.ValidatorRequest(createRequest); msgErr != nil {
		msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if isSuccess, err := e.ExpeditionService.AddExpedition(c.Request.Context(), createRequest); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if isSuccess {
			c.JSON(http.StatusOK, web.WebResponse{
				Code:    http.StatusOK,
				Status:  "OK",
				Message: "Success add expedition.",
				Data:    nil,
			})
			return
		}
	}
}

func (e *ExpeditionControllerImpl) UpdateExpedition(c *gin.Context) {
	if ID, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg("Invalid ID", exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	} else {
		var updateRequest *request.UpdateExpeditionRequest
		if err := c.ShouldBindJSON(&updateRequest); err != nil {
			msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
			exception.ErrorHandler(c, msg)
			return
		} else {

			if msgErr := helpers.ValidatorRequest(updateRequest); msgErr != nil {
				msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
				exception.ErrorHandler(c, msg)
				return
			}

			updateRequest.ID = ID
			if isSuccess, err := e.ExpeditionService.UpdateExpedition(c.Request.Context(), updateRequest); err != nil {
				msg := exception.ToErrorMsg(err.Msg, err.Error)
				exception.ErrorHandler(c, msg)
				return
			} else {
				if isSuccess {
					c.JSON(http.StatusOK, web.WebResponse{
						Code:    http.StatusOK,
						Status:  "OK",
						Message: "Success update expedition.",
						Data:    nil,
					})
					return
				}
			}
		}
	}
}

func (e *ExpeditionControllerImpl) DeleteExpedition(c *gin.Context) {
	if ID, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg("Invalid ID", exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if isSuccess, err := e.ExpeditionService.DeleteExpedition(c.Request.Context(), ID); err != nil {
			msg := exception.ToErrorMsg(err.Msg, err.Error)
			exception.ErrorHandler(c, msg)
			return
		} else {
			if isSuccess == true {
				c.JSON(http.StatusOK, web.WebResponse{
					Code:    http.StatusOK,
					Status:  "OK",
					Message: "Success delete expedition.",
					Data:    nil,
				})
				return
			}
		}
	}
}

func (e *ExpeditionControllerImpl) FindAllExpedition(c *gin.Context) {
	if expeditionResponses, err := e.ExpeditionService.FindAllExpedition(c.Request.Context()); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		c.JSON(http.StatusOK, web.WebResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "Success get all expeditions.",
			Data:    expeditionResponses,
		})
	}
}

func (e *ExpeditionControllerImpl) FindExpeditionByID(c *gin.Context) {
	if ID, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg("Invalid ID", exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if findExpeditionByID, err := e.ExpeditionService.FindExpeditionByID(c.Request.Context(), ID); err != nil {
			msg := exception.ToErrorMsg(err.Msg, err.Error)
			exception.ErrorHandler(c, msg)
			return
		} else {
			c.JSON(http.StatusOK, web.WebResponse{
				Code:    http.StatusOK,
				Status:  "OK",
				Message: "Success get expeditions.",
				Data:    findExpeditionByID,
			})
		}
	}
}
