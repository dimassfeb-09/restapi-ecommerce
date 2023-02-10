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

type AuthService interface {
	AuthLogin(ctx context.Context, request *request.AuthLoginRequest) (*response.AuthLoginResponse, *exception.ErrorMsg)
	AuthRegister(ctx context.Context, request *request.AuthRegisterRequest) (bool, *exception.ErrorMsg)
}

type AuthServiceImpl struct {
	AuthRepository repository.AuthRepository
	UserService    UserServices
	DB             *sql.DB
}

func NewAuthServiceImpl(DB *sql.DB, authRepository repository.AuthRepository, userRepository repository.UserRepository) AuthService {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: DB, UserService: &UserServiceImpl{
		DB:             DB,
		UserRepository: userRepository,
	}}
}

func (a *AuthServiceImpl) AuthLogin(ctx context.Context, request *request.AuthLoginRequest) (*response.AuthLoginResponse, *exception.ErrorMsg) {

	user := &domain.AuthUser{
		Username: request.Username,
		Password: request.Password,
	}

	if authUser, err := a.AuthRepository.AuthLogin(ctx, a.DB, user); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.BadRequest)
	} else {
		authResponse := response.ToAuthLoginResponse(authUser)
		return authResponse, nil
	}
}

func (a *AuthServiceImpl) AuthRegister(ctx context.Context, request *request.AuthRegisterRequest) (bool, *exception.ErrorMsg) {
	tx, err := a.DB.Begin()
	if err != nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if isRegistered, _ := a.UserService.FindByUsername(ctx, request.Username); isRegistered == true {
		return false, exception.ToErrorMsg("The username is already in use.", exception.BadRequest)
	}

	user := &domain.AuthUser{
		Name:     request.Name,
		Username: request.Username,
		Password: request.Password,
	}

	if isRegistered, err := a.AuthRepository.AuthRegister(ctx, tx, user); isRegistered == false {
		return false, exception.ToErrorMsg(err.Error(), exception.InternalServerError)
	}

	return true, nil
}
