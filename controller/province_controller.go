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

type ProvinceController interface {
	CreateProvince(c *gin.Context)
	UpdateProvince(c *gin.Context)
	FindByIdProvince(c *gin.Context)
	DeleteProvince(c *gin.Context)
	FindAllProvince(c *gin.Context)
}

type ProvinceControllerImpl struct {
	ProvinceService services.ProvinceService
}

func NewProvinceControllerImpl(provinceService services.ProvinceService) ProvinceController {
	return &ProvinceControllerImpl{ProvinceService: provinceService}
}

func (p *ProvinceControllerImpl) CreateProvince(c *gin.Context) {

	var createProvinceRequest *request.CreateProvinceRequest
	if err := c.ShouldBind(&createProvinceRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if msgErr := helpers.ValidatorRequest(createProvinceRequest); msgErr != nil {
		msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if response, err := p.ProvinceService.CreateProvince(c.Request.Context(), createProvinceRequest); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		webResponse := helpers.ToWebResponse(http.StatusOK, "OK", "Success Create Data.", response)
		c.JSON(http.StatusOK, webResponse)
		return
	}
}

func (p *ProvinceControllerImpl) UpdateProvince(c *gin.Context) {

	ID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	var updateProvinceRequest *request.UpdateProvinceRequest
	if err := c.ShouldBind(&updateProvinceRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}
	updateProvinceRequest.ID = ID

	if msgErr := helpers.ValidatorRequest(updateProvinceRequest); msgErr != nil {
		msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if response, err := p.ProvinceService.UpdateProvince(c.Request.Context(), updateProvinceRequest); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		webResponse := helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Update data with ID-%d", response.ID), response)
		c.JSON(http.StatusOK, webResponse)
		return
	}
}

func (p *ProvinceControllerImpl) FindByIdProvince(c *gin.Context) {

	if Id, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if response, err := p.ProvinceService.FindProvinceById(c.Request.Context(), Id); err != nil {
			msg := exception.ToErrorMsg(err.Msg, err.Error)
			exception.ErrorHandler(c, msg)
			return
		} else {
			webResponse := helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success find data with ID-%d", response.ID), response)
			c.JSON(http.StatusOK, webResponse)
			return
		}
	}

}

func (p *ProvinceControllerImpl) DeleteProvince(c *gin.Context) {

	if ID, err := strconv.Atoi(c.Param("id")); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	} else {
		if err := p.ProvinceService.DeleteProvince(c.Request.Context(), ID); err != nil {
			msg := exception.ToErrorMsg(err.Msg, err.Error)
			exception.ErrorHandler(c, msg)
			return
		} else {
			webResponse := helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success delete data with ID-%d", ID), nil)
			c.JSON(http.StatusOK, webResponse)
			return
		}
	}

}

func (p *ProvinceControllerImpl) FindAllProvince(c *gin.Context) {
	if provinceResponses, err := p.ProvinceService.FindAllProvince(c.Request.Context()); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		webResponse := helpers.ToWebResponse(http.StatusOK, "OK", "Success get Provinces", provinceResponses)
		c.JSON(http.StatusOK, webResponse)
		return
	}
}
