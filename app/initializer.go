package app

import (
	"database/sql"
	"github.com/dimassfeb-09/restapi-ecommerce.git/controller"
	"github.com/dimassfeb-09/restapi-ecommerce.git/repository"
	"github.com/dimassfeb-09/restapi-ecommerce.git/services"
)

type Controller struct {
	controller.AuthController
	controller.UserController
	controller.CityController
	controller.ProvinceController
	controller.ExpeditionController
}

func NewInitialized(db *sql.DB) *Controller {

	authRepository := repository.NewAuthRepositoryImpl()
	userRepository := repository.NewUserRepositoryImpl()
	provinceRepository := repository.NewProvinceRepositoryImpl()
	cityRepository := repository.NewCityRepositoryImpl()
	expeditionRepository := repository.NewExpeditionRepositoryImpl()

	authServices := services.NewAuthServiceImpl(db, authRepository, userRepository)
	userServices := services.NewUserServiceImpl(db, userRepository)
	provinceServices := services.NewProvinceService(db, provinceRepository, cityRepository)
	cityServices := services.NewCityServicesImpl(db, cityRepository, provinceRepository)
	expeditionServices := services.NewExpeditionServiceImpl(db, expeditionRepository)

	authController := controller.NewAuthControllerImpl(authServices)
	userController := controller.NewUserControllerImpl(userServices)
	provinceController := controller.NewProvinceControllerImpl(provinceServices)
	cityController := controller.NewCityControllerImpl(cityServices)
	expeditionController := controller.NewExpeditionControllerImpl(expeditionServices)

	return &Controller{
		AuthController:       authController,
		UserController:       userController,
		CityController:       cityController,
		ProvinceController:   provinceController,
		ExpeditionController: expeditionController,
	}
}
