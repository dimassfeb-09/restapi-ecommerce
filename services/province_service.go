package services

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
	"github.com/dimassfeb-09/restapi-ecommerce.git/helpers"
	"github.com/dimassfeb-09/restapi-ecommerce.git/repository"
)

type ProvinceService interface {
	CreateProvince(ctx context.Context, request *request.ProvinceRequest) (*response.ProvinceResponse, error)
	UpdateProvince(ctx context.Context, request *request.ProvinceRequest) (*response.ProvinceResponse, error)
	DeleteProvince(ctx context.Context, provinceId int) error
	FindProvinceById(ctx context.Context, provinceId int) (*response.ProvinceResponse, error)
	FindAllProvince(ctx context.Context) ([]*response.ProvinceResponse, error)
}

type ProvinceServiceImpl struct {
	DB                 *sql.DB
	ProvinceRepository repository.ProvinceRepository
}

func NewProvinceService(DB *sql.DB, provinceRepository repository.ProvinceRepository) ProvinceService {
	return &ProvinceServiceImpl{DB: DB, ProvinceRepository: provinceRepository}
}

func (p *ProvinceServiceImpl) CreateProvince(ctx context.Context, req *request.ProvinceRequest) (*response.ProvinceResponse, error) {

	if tx, err := p.DB.Begin(); err != nil {
		return nil, err
	} else {
		defer helpers.RollbackCommit(tx)
		if createProvince, err := p.ProvinceRepository.CreateProvince(ctx, tx, req.Name); err != nil {
			return nil, err
		} else {
			return helpers.ToProvinceResponse(createProvince), nil
		}
	}

}

func (p *ProvinceServiceImpl) UpdateProvince(ctx context.Context, req *request.ProvinceRequest) (*response.ProvinceResponse, error) {
	if tx, err := p.DB.Begin(); err != nil {
		return nil, err
	} else {
		province := &domain.Province{
			ID:   req.ID,
			Name: req.Name,
		}

		if updateProvince, err := p.ProvinceRepository.UpdateProvince(ctx, tx, province); err != nil {
			return nil, err
		} else {
			return helpers.ToProvinceResponse(updateProvince), nil
		}
	}
}

func (p *ProvinceServiceImpl) DeleteProvince(ctx context.Context, provinceId int) error {
	if tx, err := p.DB.Begin(); err != nil {
		return err
	} else {
		defer helpers.RollbackCommit(tx)
		if _, err := p.ProvinceRepository.FindProvinceById(ctx, p.DB, provinceId); err != nil {
			return exception.ErrorNotFound
		} else {
			if err := p.ProvinceRepository.DeleteProvince(ctx, tx, provinceId); err != nil {
				return err
			}
		}
		return nil
	}
}

func (p *ProvinceServiceImpl) FindProvinceById(ctx context.Context, provinceId int) (*response.ProvinceResponse, error) {
	if provinceById, err := p.ProvinceRepository.FindProvinceById(ctx, p.DB, provinceId); err != nil {
		return nil, err
	} else {
		return helpers.ToProvinceResponse(provinceById), nil
	}
}

func (p *ProvinceServiceImpl) FindAllProvince(ctx context.Context) ([]*response.ProvinceResponse, error) {
	if provinces, err := p.ProvinceRepository.FindAllProvince(ctx, p.DB); err != nil {
		return nil, err
	} else {
		return helpers.ToProvinceResponses(provinces), nil
	}
}
