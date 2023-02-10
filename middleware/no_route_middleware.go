package middleware

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/gin-gonic/gin"
)

func NoRoute(c *gin.Context) {
	c.JSON(404, helpers.ToErrorResponse(404, "PAGE_NOT_FOUND", "PAGE NOT FOUND"))
	return
}
