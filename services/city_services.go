package services

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/dimassfeb-09/restapi-ecommerce.git/repository"
)

type CityServices interface {
	CreateCity(ctx context.Context, city *request.CreateCityRequest) (*response.CityResponse, *exception.ErrorMsg)
	FindByIdCity(ctx context.Context, cityId int) (*response.CityResponse, *exception.ErrorMsg)
	UpdateCity(ctx context.Context, updateReq *request.UpdateCityRequest) (*response.CityResponse, *exception.ErrorMsg)
	DeleteCity(ctx context.Context, cityId int) *exception.ErrorMsg
	FindAllCity(ctx context.Context) ([]*response.CityResponse, *exception.ErrorMsg)
}

type CityServicesImpl struct {
	DB             *sql.DB
	CityRepository repository.CityRepository
}

func NewCityServicesImpl(DB *sql.DB, cityRepository repository.CityRepository) *CityServicesImpl {
	return &CityServicesImpl{CityRepository: cityRepository, DB: DB}
}

func (c *CityServicesImpl) CreateCity(ctx context.Context, cityReqCreate *request.CreateCityRequest) (*response.CityResponse, *exception.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	}
	defer helpers.RollbackCommit(tx)

	if findCityName, _ := c.CityRepository.FindCityByName(ctx, c.DB, cityReqCreate.Name); findCityName.Name == cityReqCreate.Name {
		msg := fmt.Sprintf("City with Name %s already exists.", cityReqCreate.Name)
		return nil, exception.ToErrorMsg(msg, exception.BadRequest)
	}

	if city, err := c.CityRepository.CreateCity(ctx, tx, cityReqCreate.Name); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		return helpers.ToCityResponse(city), nil
	}
}

func (c *CityServicesImpl) FindByIdCity(ctx context.Context, cityId int) (*response.CityResponse, *exception.ErrorMsg) {
	if city, err := c.CityRepository.FindCityById(ctx, c.DB, cityId); err != nil {
		msg := fmt.Sprintf("City with ID-%d Not Found.", cityId)
		return nil, exception.ToErrorMsg(msg, err)
	} else {
		return helpers.ToCityResponse(city), nil
	}
}

func (c *CityServicesImpl) UpdateCity(ctx context.Context, updateReq *request.UpdateCityRequest) (*response.CityResponse, *exception.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if _, err := c.CityRepository.FindCityById(ctx, c.DB, updateReq.ID); err != nil {
		msg := fmt.Sprintf("Data with ID-%d Not Found", updateReq.ID)
		return nil, exception.ToErrorMsg(msg, exception.BadRequest)
	}

	if findByName, _ := c.CityRepository.FindCityByName(ctx, c.DB, updateReq.Name); findByName != nil {
		if findByName.ID != updateReq.ID {
			msg := fmt.Sprintf("City with Name %s already exists.", updateReq.Name)
			return nil, exception.ToErrorMsg(msg, exception.BadRequest)
		} else if findByName.Name == updateReq.Name {
			msg := fmt.Sprintf("Not allowed with the previous name.")
			return nil, exception.ToErrorMsg(msg, exception.BadRequest)
		}
	}

	city := &domain.City{
		ID:   updateReq.ID,
		Name: updateReq.Name,
	}

	if updateCity, err := c.CityRepository.UpdateCity(ctx, tx, city); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		return helpers.ToCityResponse(updateCity), nil
	}
}

func (c *CityServicesImpl) DeleteCity(ctx context.Context, cityId int) *exception.ErrorMsg {
	tx, err := c.DB.Begin()
	if err != nil {
		return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	}
	defer helpers.RollbackCommit(tx)

	if cityById, _ := c.CityRepository.FindCityById(ctx, c.DB, cityId); cityById != nil {
		msg := fmt.Sprintf("Kota dengan ID %d tidak ditemukan", cityId)
		return exception.ToErrorMsg(msg, exception.BadRequest)
	} else {
		if err := c.CityRepository.DeleteCity(ctx, tx, cityId); err != nil {
			return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
		}
		return nil
	}

}

func (c *CityServicesImpl) FindAllCity(ctx context.Context) ([]*response.CityResponse, *exception.ErrorMsg) {
	if cities, err := c.CityRepository.FindAllCity(ctx, c.DB); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		return helpers.ToCityResponses(cities), nil
	}
}
