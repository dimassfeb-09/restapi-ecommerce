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
	"time"
)

type UserServices interface {
	CreateUser(ctx context.Context, users *request.CreateUserRequest) (*response.UserResponse, *exception.ErrorMsg)
	FindByIdUser(ctx context.Context, userId int) (*response.UserResponse, *exception.ErrorMsg)
	UpdateUser(ctx context.Context, updateReq *request.UpdateUserRequest) (*response.UserResponse, *exception.ErrorMsg)
	DeleteUser(ctx context.Context, userId int) *exception.ErrorMsg
	FindAllUser(ctx context.Context) ([]*response.UserResponse, *exception.ErrorMsg)
}

type UserServiceImpl struct {
	DB             *sql.DB
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(DB *sql.DB, userRepository repository.UserRepository) UserServices {
	return &UserServiceImpl{DB: DB, UserRepository: userRepository}
}

func (u *UserServiceImpl) CreateUser(ctx context.Context, req *request.CreateUserRequest) (*response.UserResponse, *exception.ErrorMsg) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	// Check Username Data Registered
	if isRegistered, _ := u.UserRepository.FindByUsername(ctx, u.DB, req.Username); isRegistered == true {
		errMsg := fmt.Sprintf("Username dengan %s sudah digunakan.", req.Username)
		return nil, exception.ToErrorMsg(errMsg, exception.BadRequest)
	}

	user := &domain.Users{
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
	}

	if createUser, err := u.UserRepository.CreateUser(ctx, tx, user); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		createUser.CreatedAt = time.Now()
		return helpers.ToUserResponse(createUser), nil
	}
}

func (u *UserServiceImpl) FindByIdUser(ctx context.Context, userId int) (*response.UserResponse, *exception.ErrorMsg) {
	db := u.DB
	if user, err := u.UserRepository.FindByIdUser(ctx, db, userId); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.ErrorNotFound)
	} else {
		return helpers.ToUserResponse(user), nil
	}
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, updateReq *request.UpdateUserRequest) (*response.UserResponse, *exception.ErrorMsg) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if userById, err := u.UserRepository.FindByIdUser(ctx, u.DB, updateReq.ID); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.ErrorNotFound)
	} else if updateReq.Username != userById.Username {
		if isRegistered, _ := u.UserRepository.FindByUsername(ctx, u.DB, updateReq.Username); isRegistered == true {
			errMsg := fmt.Sprintf("Username dengan %s sudah digunakan.", updateReq.Username)
			return nil, exception.ToErrorMsg(errMsg, exception.BadRequest)
		}
	}

	user := &domain.Users{
		ID:       updateReq.ID,
		Name:     updateReq.Name,
		Username: updateReq.Username,
		Password: updateReq.Password,
		Balance:  updateReq.Balance,
	}

	if updateUser, err := u.UserRepository.UpdateUser(ctx, tx, user); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		return helpers.ToUserResponse(updateUser), nil
	}
}

func (u *UserServiceImpl) DeleteUser(ctx context.Context, userId int) *exception.ErrorMsg {
	tx, err := u.DB.Begin()
	if err != nil {
		return exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if _, err := u.UserRepository.FindByIdUser(ctx, u.DB, userId); err != nil {
		return exception.ToErrorMsg(err.Error(), exception.ErrorNotFound)
	} else {
		if err := u.UserRepository.DeleteUser(ctx, tx, userId); err != nil {
			return exception.ToErrorMsg(err.Error(), exception.InternalServerError)
		}
	}

	return nil
}

func (u *UserServiceImpl) FindAllUser(ctx context.Context) ([]*response.UserResponse, *exception.ErrorMsg) {

	if user, err := u.UserRepository.FindAll(ctx, u.DB); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		userResponses := helpers.ToUserResponses(user)
		return userResponses, nil
	}
}
