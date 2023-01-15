package exception

import (
	"errors"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/web"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrorNotFound       error = errors.New("Not Found")
	BadRequest          error = errors.New("Bad Request")
	InternalServerError error = errors.New("Internal Server Error")
)

type ErrorMsg struct {
	Msg   string
	Error error
}

func ToErrorMsg(msg string, err error) *ErrorMsg {
	return &ErrorMsg{
		Msg:   msg,
		Error: err,
	}
}

func ToErrorResponse(code int, status string, msg string) *web.ErrorResponse {
	return &web.ErrorResponse{
		Code:    code,
		Status:  status,
		Message: msg,
	}
}

func ErrorHandler(c *gin.Context, err *ErrorMsg) {
	var errHandler *web.ErrorResponse
	switch err.Error {
	case ErrorNotFound:
		errHandler = ToErrorResponse(http.StatusNotFound, ErrorNotFound.Error(), err.Msg)
		break
	case BadRequest:
		errHandler = ToErrorResponse(http.StatusBadRequest, BadRequest.Error(), err.Msg)
		break
	default:
		errHandler = ToErrorResponse(http.StatusInternalServerError, InternalServerError.Error(), err.Msg)
		break
	}

	c.AbortWithStatusJSON(errHandler.Code, helpers.ToErrorResponse(errHandler.Code, errHandler.Status, errHandler.Message))
}
