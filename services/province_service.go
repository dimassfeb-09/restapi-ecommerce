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
	"log"
)

type ProvinceService interface {
	CreateProvince(ctx context.Context, request *request.CreateProvinceRequest) (*response.ProvinceResponse, *exception.ErrorMsg)
	UpdateProvince(ctx context.Context, request *request.UpdateProvinceRequest) (*response.ProvinceResponse, *exception.ErrorMsg)
	DeleteProvince(ctx context.Context, provinceId int) *exception.ErrorMsg
	FindProvinceById(ctx context.Context, provinceId int) (*response.ProvinceResponse, *exception.ErrorMsg)
	FindProvinceByName(ctx context.Context, name string) (*response.ProvinceResponse, *exception.ErrorMsg)
	FindAllProvince(ctx context.Context) ([]*response.ProvinceResponse, *exception.ErrorMsg)
}

type ProvinceServiceImpl struct {
	DB                 *sql.DB
	ProvinceRepository repository.ProvinceRepository
	CityServices       CityServices
}

func NewProvinceService(DB *sql.DB, provinceRepository repository.ProvinceRepository, cityRepository repository.CityRepository) ProvinceService {
	return &ProvinceServiceImpl{DB: DB, ProvinceRepository: provinceRepository, CityServices: &CityServicesImpl{
		DB:             DB,
		CityRepository: cityRepository,
	}}
}

func (p *ProvinceServiceImpl) CreateProvince(ctx context.Context, req *request.CreateProvinceRequest) (*response.ProvinceResponse, *exception.ErrorMsg) {

	if tx, err := p.DB.Begin(); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		defer helpers.RollbackCommit(tx)

		if len(req.Name) < 3 || len(req.Name) > 20 {
			return nil, exception.ToErrorMsg("Min length name 3 character, Max 20 character.", exception.BadRequest)
		}

		if findProvinceByName, _ := p.ProvinceRepository.FindProvinceByName(ctx, p.DB, req.Name); findProvinceByName != nil {
			if findProvinceByName.Name == req.Name {
				msg := fmt.Sprintf("City with Name %s already exists.", req.Name)
				return nil, exception.ToErrorMsg(msg, exception.BadRequest)
			}
		}

		if createProvince, err := p.ProvinceRepository.CreateProvince(ctx, tx, req.Name); err != nil {
			return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
		} else {
			return response.ToProvinceResponse(createProvince), nil
		}
	}

}

func (p *ProvinceServiceImpl) UpdateProvince(ctx context.Context, req *request.UpdateProvinceRequest) (*response.ProvinceResponse, *exception.ErrorMsg) {
	if tx, err := p.DB.Begin(); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		defer helpers.RollbackCommit(tx)

		if len(req.Name) < 3 || len(req.Name) > 20 {
			return nil, exception.ToErrorMsg("Min length name 3 character, Max 20 character.", exception.BadRequest)
		}

		province := &domain.Province{
			ID:   req.ID,
			Name: req.Name,
		}

		if findByID, _ := p.ProvinceRepository.FindProvinceById(ctx, p.DB, req.ID); findByID == nil {
			msg := fmt.Sprintf("Province with ID-%d not found.", req.ID)
			return nil, exception.ToErrorMsg(msg, exception.BadRequest)
		}

		if findProvinceByName, _ := p.ProvinceRepository.FindProvinceByName(ctx, p.DB, req.Name); findProvinceByName != nil {
			if findProvinceByName.Name == req.Name {
				msg := fmt.Sprintf("City with Name %s already exists.", req.Name)
				return nil, exception.ToErrorMsg(msg, exception.BadRequest)
			}
		}

		if updateProvince, errs := p.ProvinceRepository.UpdateProvince(ctx, tx, province); err != nil {
			return nil, exception.ToErrorMsg(errs.Error(), exception.InternalServerError)
		} else {
			return response.ToProvinceResponse(updateProvince), nil
		}
	}
}

func (p *ProvinceServiceImpl) DeleteProvince(ctx context.Context, provinceId int) *exception.ErrorMsg {
	if tx, err := p.DB.Begin(); err != nil {
		return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		defer helpers.RollbackCommit(tx)
		if _, err := p.ProvinceRepository.FindProvinceById(ctx, p.DB, provinceId); err != nil {
			return exception.ToErrorMsg(fmt.Sprintf("Province with ID-%d not found.", provinceId), exception.ErrorNotFound)
		}

		if findCityByProvID, _ := p.CityServices.FindCityByProvinceID(ctx, provinceId); findCityByProvID == true {
			return exception.ToErrorMsg("Can't delete data, because it is related to the city.", exception.BadRequest)
		}

		if err := p.ProvinceRepository.DeleteProvince(ctx, tx, provinceId); err != nil {
			return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
		}

		return nil
	}
}

func (p *ProvinceServiceImpl) FindProvinceById(ctx context.Context, provinceId int) (*response.ProvinceResponse, *exception.ErrorMsg) {
	if provinceById, err := p.ProvinceRepository.FindProvinceById(ctx, p.DB, provinceId); err != nil {
		return nil, exception.ToErrorMsg(fmt.Sprintf("Province with ID-%d not found", provinceId), exception.BadRequest)
	} else {
		return response.ToProvinceResponse(provinceById), nil
	}
}

func (p *ProvinceServiceImpl) FindProvinceByName(ctx context.Context, name string) (*response.ProvinceResponse, *exception.ErrorMsg) {
	if findByName, err := p.ProvinceRepository.FindProvinceByName(ctx, p.DB, name); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		return response.ToProvinceResponse(findByName), nil
	}
}

func (p *ProvinceServiceImpl) FindAllProvince(ctx context.Context) ([]*response.ProvinceResponse, *exception.ErrorMsg) {
	if provinces, err := p.ProvinceRepository.FindAllProvince(ctx, p.DB); err != nil {
		log.Println("FindAllProvince", err)
		return nil, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	} else {
		return response.ToProvinceResponses(provinces), nil
	}
}
