package city

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
	"github.com/dimassfeb-09/restapi-ecommerce.git/usecases/province"
)

type CityServices interface {
	CreateCity(ctx context.Context, city *request.CreateCityRequest) (*response.CityResponse, *exception.ErrorMsg)
	UpdateCity(ctx context.Context, updateReq *request.UpdateCityRequest) (*response.CityResponse, *exception.ErrorMsg)
	DeleteCity(ctx context.Context, cityId int) *exception.ErrorMsg
	FindCityByID(ctx context.Context, cityId int) (*response.CityResponse, *exception.ErrorMsg)
	FindCityByProvinceID(ctx context.Context, provinceId int) (bool, *exception.ErrorMsg)
	FindAllCity(ctx context.Context) ([]*response.CityResponse, *exception.ErrorMsg)
}

type CityServicesImpl struct {
	DB              *sql.DB
	CityRepository  repository.CityRepository
	ProvinceService province.ProvinceService
}

func NewCityServicesImpl(DB *sql.DB, cityRepository repository.CityRepository, provinceRepository repository.ProvinceRepository) CityServices {
	return &CityServicesImpl{DB: DB, CityRepository: cityRepository, ProvinceService: &province.ProvinceServiceImpl{
		DB:                 DB,
		ProvinceRepository: provinceRepository,
	}}
}

func (c *CityServicesImpl) CreateCity(ctx context.Context, createReq *request.CreateCityRequest) (*response.CityResponse, *exception.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	}
	defer helpers.RollbackCommit(tx)

	if len(createReq.Name) < 3 || len(createReq.Name) > 20 {
		return nil, exception.ToErrorMsg("Min length name 3 character, Max 20 character.", exception.BadRequest)
	}

	if findCityByName, _ := c.CityRepository.FindCityByName(ctx, c.DB, createReq.Name); findCityByName != nil {
		msg := fmt.Sprintf("City with name %s already exists.", createReq.Name)
		return nil, exception.ToErrorMsg(msg, exception.BadRequest)
	}

	if findProvincebyId, _ := c.ProvinceService.FindProvinceById(ctx, createReq.ProvinceID); findProvincebyId == nil {
		msg := fmt.Sprintf("province_id with ID-%d Not Found.", createReq.ProvinceID)
		return nil, exception.ToErrorMsg(msg, exception.BadRequest)
	}

	if city, err := c.CityRepository.CreateCity(ctx, tx, createReq.Name, createReq.ProvinceID); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		return response.ToCityResponse(city), nil
	}
}

func (c *CityServicesImpl) UpdateCity(ctx context.Context, updateReq *request.UpdateCityRequest) (*response.CityResponse, *exception.ErrorMsg) {
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if len(updateReq.Name) < 3 || len(updateReq.Name) > 20 {
		return nil, exception.ToErrorMsg("Min length name 3 character, Max 20 character.", exception.BadRequest)
	}

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

	//if _, errProv := c.ProvinceService.FindProvinceById(ctx, updateReq.ProvinceID); errProv != nil {
	//	msg := fmt.Sprintf("province_id with ID-%d Not Found.", updateReq.ProvinceID)
	//	return nil, exception.ToErrorMsg(msg, exception.BadRequest)
	//}

	city := &domain.City{
		ID:         updateReq.ID,
		Name:       updateReq.Name,
		ProvinceID: updateReq.ProvinceID,
	}

	if updateCity, err := c.CityRepository.UpdateCity(ctx, tx, city); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		return response.ToCityResponse(updateCity), nil
	}
}

func (c *CityServicesImpl) DeleteCity(ctx context.Context, cityId int) *exception.ErrorMsg {
	tx, err := c.DB.Begin()
	if err != nil {
		return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	}
	defer helpers.RollbackCommit(tx)

	if cityById, _ := c.CityRepository.FindCityById(ctx, c.DB, cityId); cityById == nil {
		msg := fmt.Sprintf("Kota dengan ID %d tidak ditemukan", cityId)
		return exception.ToErrorMsg(msg, exception.BadRequest)
	} else {
		if err := c.CityRepository.DeleteCity(ctx, tx, cityId); err != nil {
			return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
		}
		return nil
	}

}

func (c *CityServicesImpl) FindCityByID(ctx context.Context, cityId int) (*response.CityResponse, *exception.ErrorMsg) {
	if city, err := c.CityRepository.FindCityById(ctx, c.DB, cityId); err != nil {
		msg := fmt.Sprintf("City with ID-%d Not Found.", cityId)
		return nil, exception.ToErrorMsg(msg, err)
	} else {
		return response.ToCityResponse(city), nil
	}
}

func (c *CityServicesImpl) FindCityByProvinceID(ctx context.Context, provinceId int) (bool, *exception.ErrorMsg) {
	isValid, err := c.CityRepository.FindCityByProvinceID(ctx, c.DB, provinceId)
	if isValid {
		return true, nil
	}

	return false, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
}

func (c *CityServicesImpl) FindAllCity(ctx context.Context) ([]*response.CityResponse, *exception.ErrorMsg) {
	if cities, err := c.CityRepository.FindAllCity(ctx, c.DB); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		return response.ToCityResponses(cities), nil
	}
}
