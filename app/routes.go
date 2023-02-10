package app

import (
	"github.com/dimassfeb-09/restapi-ecommerce.git/middleware"
	"github.com/gin-gonic/gin"
)

func RoutesApi(initializer *Controller, r *gin.Engine) {

	api := r.Group("api")
	auth := initializer.AuthController
	api.POST("/auth/login", auth.AuthLogin)
	api.POST("/auth/register", auth.AuthRegister)

	api.Use(middleware.JWTMiddleware)
	user := initializer.UserController
	api.GET("/user/:id", user.FindByIdUser)
	api.GET("/users", user.FindAllUser)
	api.POST("/user", user.CreateUser)
	api.PUT("/user/:id", user.UpdateUser)
	api.DELETE("/user/:id", user.DeleteUser)
	api.PUT("/user/update/password/:id", user.ChangePassword)

	province := initializer.ProvinceController
	api.POST("/province", province.CreateProvince)
	api.GET("/province/:id", province.FindByIdProvince)
	api.GET("/provinces", province.FindAllProvince)
	api.PUT("/province/:id", province.UpdateProvince)
	api.DELETE("/province/:id", province.DeleteProvince)

	city := initializer.CityController
	api.GET("/city/:id", city.FindByIdCity)
	api.GET("/cities", city.FindAllCity)
	api.POST("/city", city.CreateCity)
	api.PUT("/city/:id", city.UpdateCity)
	api.DELETE("/city/:id", city.DeleteCity)

	exp := initializer.ExpeditionController
	api.GET("/expedition/:id", exp.FindExpeditionByID)
	api.GET("/expeditions", exp.FindAllExpedition)
	api.POST("/expedition", exp.AddExpedition)
	api.PUT("/expedition/:id", exp.UpdateExpedition)
	api.DELETE("/expedition/:id", exp.DeleteExpedition)

}
