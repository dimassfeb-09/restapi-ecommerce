package controller

import (
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/dimassfeb-09/restapi-ecommerce.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type CityController interface {
	CreateCity(c *gin.Context)
	UpdateCity(c *gin.Context)
	FindByIdCity(c *gin.Context)
	DeleteCity(c *gin.Context)
	FindAllCity(c *gin.Context)
}

type CityControllerImpl struct {
	CityService services.CityServices
}

func (cityC *CityControllerImpl) CreateCity(c *gin.Context) {

	var createCityRequest *request.CreateCityRequest
	if err := c.ShouldBind(&createCityRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if city, errMsg := cityC.CityService.CreateCity(c.Request.Context(), createCityRequest); errMsg != nil {
		msg := exception.ToErrorMsg(errMsg.Msg, errMsg.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Created Data with ID-%d", city.ID), city))
		return
	}
}

func (cityC *CityControllerImpl) UpdateCity(c *gin.Context) {
	var updateCityRequest *request.UpdateCityRequest

	Id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if err := c.ShouldBind(&updateCityRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	updateCityRequest.ID = Id
	if city, errMsg := cityC.CityService.UpdateCity(c.Request.Context(), updateCityRequest); errMsg != nil {
		msg := exception.ToErrorMsg(errMsg.Msg, errMsg.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Update City with ID-%d", city.ID), city))
		return
	}
}

func (cityC *CityControllerImpl) FindByIdCity(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if city, errMsg := cityC.CityService.FindCityByID(c.Request.Context(), id); errMsg != nil {
			msg := exception.ToErrorMsg(errMsg.Msg, errMsg.Error)
			exception.ErrorHandler(c, msg)
			return
		} else {
			c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Get Data with ID-%d", id), city))
			return
		}
	}
}

func (cityC *CityControllerImpl) DeleteCity(c *gin.Context) {
	if id, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg(err.Error(), err)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if errMsg := cityC.CityService.DeleteCity(c.Request.Context(), id); errMsg != nil {
			msg := exception.ToErrorMsg(errMsg.Msg, errMsg.Error)
			exception.ErrorHandler(c, msg)
			return
		} else {
			c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Deleted Data with ID-%d", id), ""))
			return
		}
	}
}

func (cityC *CityControllerImpl) FindAllCity(c *gin.Context) {
	if city, errMsg := cityC.CityService.FindAllCity(c.Request.Context()); errMsg != nil {
		msg := exception.ToErrorMsg(errMsg.Msg, errMsg.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", "Success Get Cities.", city))
		return
	}
}

func NewCityControllerImpl(cityService services.CityServices) CityController {
	return &CityControllerImpl{CityService: cityService}
}
