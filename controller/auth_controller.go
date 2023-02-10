package controller

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/dimassfeb-09/restapi-ecommerce.git/middleware"
	"github.com/dimassfeb-09/restapi-ecommerce.git/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type AuthController interface {
	AuthLogin(c *gin.Context)
	AuthRegister(c *gin.Context)
}

type AuthControllerImpl struct {
	AuthService services.AuthService
}

func NewAuthControllerImpl(authService services.AuthService) AuthController {
	return &AuthControllerImpl{AuthService: authService}
}

func (a *AuthControllerImpl) AuthLogin(c *gin.Context) {

	var authLoginRequest request.AuthLoginRequest
	if err := c.ShouldBindJSON(&authLoginRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if authLoginResponse, err := a.AuthService.AuthLogin(c.Request.Context(), &authLoginRequest); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else if authLoginResponse != nil {
		user := &middleware.User{
			ID:       authLoginResponse.ID,
			Username: authLoginRequest.Username,
		}
		if token, err := middleware.NewGenerateJWTToken(user); err != nil {
			msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
			exception.ErrorHandler(c, msg)
			return
		} else {
			token = strings.Trim(token, "\"")
			authLoginResponse.Token = token
			c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", "Success Login.", authLoginResponse))
			return
		}
	}
}

func (a *AuthControllerImpl) AuthRegister(c *gin.Context) {
	var authRegisterRequest *request.AuthRegisterRequest
	if err := c.ShouldBindJSON(&authRegisterRequest); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if isSuccess, err := a.AuthService.AuthRegister(c.Request.Context(), authRegisterRequest); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else if isSuccess {
		c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", "Success Register.", nil))
		return
	}

}
