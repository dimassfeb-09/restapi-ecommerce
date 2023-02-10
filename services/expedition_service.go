package expedition

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

type ExpeditionService interface {
	AddExpedition(ctx context.Context, request *request.ExpeditionCreateRequest) (bool, *exception.ErrorMsg)
	UpdateExpedition(ctx context.Context, updateRequest *request.ExpeditionUpdateRequest) (bool, *exception.ErrorMsg)
	DeleteExpedition(ctx context.Context, expID int) (bool, *exception.ErrorMsg)
	FindAllExpedition(ctx context.Context) ([]*response.ExpeditionResponse, *exception.ErrorMsg)
	FindExpeditionByID(ctx context.Context, expID int) (*response.ExpeditionResponse, *exception.ErrorMsg)
}

type ExpeditionServiceImpl struct {
	DB                   *sql.DB
	ExpeditionRepository repository.ExpeditionRepository
}

func NewExpeditionServiceImpl(DB *sql.DB, expeditionRepository repository.ExpeditionRepository) ExpeditionService {
	return &ExpeditionServiceImpl{DB: DB, ExpeditionRepository: expeditionRepository}
}

func (e *ExpeditionServiceImpl) AddExpedition(ctx context.Context, request *request.ExpeditionCreateRequest) (bool, *exception.ErrorMsg) {
	tx, err := e.DB.Begin()
	if err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	expedition := &domain.Expedition{
		Name: request.Name,
	}

	if isSuccess, err := e.ExpeditionRepository.AddExpedition(ctx, tx, expedition); err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	} else {
		if isSuccess {
			return true, nil
		} else {
			return false, exception.ToErrorMsg(err.Error(), err)
		}
	}
}

func (e *ExpeditionServiceImpl) UpdateExpedition(ctx context.Context, request *request.ExpeditionUpdateRequest) (bool, *exception.ErrorMsg) {
	tx, err := e.DB.Begin()
	if err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	expedition := &domain.Expedition{
		Name: request.Name,
	}

	if isSuccess, err := e.ExpeditionRepository.UpdateExpedition(ctx, tx, expedition); err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	} else {
		if isSuccess {
			return true, nil
		} else {
			return false, exception.ToErrorMsg(err.Error(), err)
		}
	}
}

func (e *ExpeditionServiceImpl) DeleteExpedition(ctx context.Context, expID int) (bool, *exception.ErrorMsg) {
	tx, err := e.DB.Begin()
	if err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if isSuccess, err := e.ExpeditionRepository.DeleteExpedition(ctx, tx, expID); err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	} else {
		if isSuccess {
			return true, nil
		} else {
			return false, exception.ToErrorMsg(err.Error(), err)
		}
	}
}

func (e *ExpeditionServiceImpl) FindAllExpedition(ctx context.Context) ([]*response.ExpeditionResponse, *exception.ErrorMsg) {
	//TODO implement me
	panic("implement me")
}

func (e *ExpeditionServiceImpl) FindExpeditionByID(ctx context.Context, expID int) (*response.ExpeditionResponse, *exception.ErrorMsg) {

	if result, err := e.ExpeditionRepository.FindExpeditionByID(ctx, e.DB, expID); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		if result != nil {
			return response.ToExpeditionResponse(result), nil
		} else {
			return nil, exception.ToErrorMsg(err.Error(), err)
		}
	}
}
