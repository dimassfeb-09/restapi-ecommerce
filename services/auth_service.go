package auth

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/request"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/response"
	"github.com/dimassfeb-09/restapi-ecommerce.git/repository"
)

type AuthService interface {
	AuthLogin(ctx context.Context, request *request.AuthLoginRequest) (*response.AuthLoginResponse, error)
}

type AuthServiceImpl struct {
	AuthRepository  repository.AuthRepository
	LoginRepository repository.UserRepository
	DB              *sql.DB
}

func NewAuthServiceImpl(DB *sql.DB, authRepository repository.AuthRepository) AuthService {
	return &AuthServiceImpl{AuthRepository: authRepository, DB: DB}
}

func (a *AuthServiceImpl) AuthLogin(ctx context.Context, request *request.AuthLoginRequest) (*response.AuthLoginResponse, error) {

	user := &domain.AuthUser{
		Username: request.Username,
		Password: request.Password,
	}

	if authUser, err := a.AuthRepository.AuthLogin(ctx, a.DB, user); err != nil {
		return nil, err
	} else {
		authResponse := response.ToAuthLoginResponse(authUser)
		return authResponse, nil
	}
}
