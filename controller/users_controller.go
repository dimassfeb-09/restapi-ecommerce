package controller

import (
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/web"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/dimassfeb-09/restapi-ecommerce.git/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"strconv"
	"strings"
)

type UserController interface {
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	FindByIdUser(c *gin.Context)
	FindByUsername(c *gin.Context)
	DeleteUser(c *gin.Context)
	FindAllUser(c *gin.Context)
	ChangePassword(c *gin.Context)
}

type UserControllerImpl struct {
	UserService services.UserServices
}

func NewUserControllerImpl(userService services.UserServices) UserController {
	return &UserControllerImpl{UserService: userService}
}

func (u *UserControllerImpl) CreateUser(c *gin.Context) {

	var user *request.CreateUserRequest
	if err := c.ShouldBind(&user); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if msgErr := helpers.ValidatorRequest(user); msgErr != nil {
		msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	validate := validator.New()
	if err := validate.Struct(user); err != nil {
		var msgErr []string
		for _, err := range err.(validator.ValidationErrors) {
			field := strings.ToLower(err.Field())
			fmt.Println(field)
			tag := err.Tag()
			msgErr = append(msgErr, "Error with field", field, "where Tag", tag)
		}
		msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if createResponse, err := u.UserService.CreateUser(c.Request.Context(), user); err != nil { // if failed
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else { // if success
		c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Created Data with ID-%d", createResponse.ID), createResponse))
		return
	}

}

func (u *UserControllerImpl) UpdateUser(c *gin.Context) {

	id := c.Param("id")
	dataId, _ := strconv.Atoi(id)

	var user *request.UpdateUserRequest
	if err := c.ShouldBind(&user); err != nil {
		msg := exception.ToErrorMsg(err.Error(), exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	if msgErr := helpers.ValidatorRequest(user); msgErr != nil {
		msg := exception.ToErrorMsg(msgErr[0], exception.BadRequest)
		exception.ErrorHandler(c, msg)
		return
	}

	user.ID = dataId
	if updateUser, err := u.UserService.UpdateUser(c.Request.Context(), user); err != nil {
		msg := exception.ToErrorMsg(err.Msg, err.Error)
		exception.ErrorHandler(c, msg)
		return
	} else {
		c.JSON(http.StatusOK, helpers.ToWebResponse(http.StatusOK, "OK", fmt.Sprintf("Success Update Data with ID-%d", user.ID), updateUser))
		return
	}

}

func (u *UserControllerImpl) FindByIdUser(c *gin.Context) {

	id := c.Param("id")
	dataId, err := strconv.Atoi(id)
	if err != nil {
		msg := fmt.Sprintf("Parameter %s tidak diizinkan. Gunakan paramter Integer.", id)
		errMsg := exception.ToErrorMsg(msg, exception.BadRequest)
		exception.ErrorHandler(c, errMsg)
		return
	}

	user, errs := u.UserService.FindByIdUser(c.Request.Context(), dataId)
	if errs != nil {
		msg := exception.ToErrorMsg(errs.Msg, errs.Error)
		exception.ErrorHandler(c, msg)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: fmt.Sprintf("Success Get Data with ID-%d", user.ID),
		Data:    user,
	})
}

func (u *UserControllerImpl) FindByUsername(c *gin.Context) {
	// implementmen
	panic("implement me")
}

func (u *UserControllerImpl) DeleteUser(c *gin.Context) {

	id := c.Param("id")
	dataId, err := strconv.Atoi(id)
	if err != nil {
		numError := err.(*strconv.NumError)
		s := numError.Err.Error()
		if s == "invalid syntax" {
			msg := fmt.Sprintf("Errors: %s bukan type data number.", numError.Num)
			errMsg := exception.ToErrorMsg(msg, exception.BadRequest)
			exception.ErrorHandler(c, errMsg)
			return
		}
	}

	if errs := u.UserService.DeleteUser(c.Request.Context(), dataId); errs != nil {
		msg := exception.ToErrorMsg(errs.Msg, errs.Error)
		exception.ErrorHandler(c, msg)
		return
	}

	c.JSON(http.StatusOK, web.WebResponse{
		Code:    http.StatusOK,
		Status:  "OK",
		Message: "Berhasil menghapus data dengan ID " + id,
	})
}

func (u *UserControllerImpl) FindAllUser(c *gin.Context) {

	if users, errMsg := u.UserService.FindAllUser(c.Request.Context()); errMsg != nil {
		err := exception.ToErrorMsg(errMsg.Msg, errMsg.Error)
		exception.ErrorHandler(c, err)
	} else {
		c.JSON(http.StatusOK, web.WebResponse{
			Code:    http.StatusOK,
			Status:  "OK",
			Message: "Success Get All Data Users",
			Data:    users,
		})
	}

}

func (u *UserControllerImpl) ChangePassword(c *gin.Context) {

	if ID, err := strconv.Atoi(c.Param("id")); err != nil {
		err := exception.ToErrorMsg("ID not valid.", exception.BadRequest)
		exception.ErrorHandler(c, err)
		return
	} else {

		var changePassword *request.ChangePasswordRequest
		if err := c.ShouldBindJSON(&changePassword); err != nil {
			err := exception.ToErrorMsg(err.Error(), exception.BadRequest)
			exception.ErrorHandler(c, err)
			return
		}

		changePassword.ID = ID
		if err := u.UserService.ChangePassword(c.Request.Context(), changePassword); err != nil {
			err := exception.ToErrorMsg(err.Msg, err.Error)
			exception.ErrorHandler(c, err)
			return
		} else {
			c.JSON(http.StatusOK, web.WebResponse{
				Code:    http.StatusOK,
				Status:  "OK",
				Message: "Success change password",
				Data:    nil,
			})
			return
		}

	}

}
