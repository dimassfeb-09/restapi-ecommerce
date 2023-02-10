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
	UpdateUser(ctx context.Context, req *request.UpdateUserRequest) (*response.UserResponse, *exception.ErrorMsg)
	DeleteUser(ctx context.Context, userId int) *exception.ErrorMsg
	FindAllUser(ctx context.Context) ([]*response.UserResponse, *exception.ErrorMsg)
	ChangePassword(ctx context.Context, req *request.ChangePasswordRequest) *exception.ErrorMsg
	FindByUsername(ctx context.Context, username string) (bool, *exception.ErrorMsg)
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

	if findByUsername, _ := u.UserRepository.FindByUsername(ctx, u.DB, req.Username); findByUsername != nil {
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
		return response.ToUserResponse(createUser), nil
	}
}

func (u *UserServiceImpl) FindByIdUser(ctx context.Context, userId int) (*response.UserResponse, *exception.ErrorMsg) {
	db := u.DB
	if user, err := u.UserRepository.FindByIdUser(ctx, db, userId); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.ErrorNotFound)
	} else {
		return response.ToUserResponse(user), nil
	}
}

func (u *UserServiceImpl) UpdateUser(ctx context.Context, req *request.UpdateUserRequest) (*response.UserResponse, *exception.ErrorMsg) {
	tx, err := u.DB.Begin()
	if err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if userById, err := u.UserRepository.FindByIdUser(ctx, u.DB, req.ID); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), exception.ErrorNotFound)
	} else if req.Username != userById.Username {
		if findByUsername, _ := u.UserRepository.FindByUsername(ctx, u.DB, req.Username); findByUsername != nil {
			errMsg := fmt.Sprintf("Username dengan %s sudah digunakan.", req.Username)
			return nil, exception.ToErrorMsg(errMsg, exception.BadRequest)
		}
	}

	user := &domain.Users{
		ID:       req.ID,
		Name:     req.Name,
		Username: req.Username,
		Password: req.Password,
		Balance:  req.Balance,
	}

	if updateUser, err := u.UserRepository.UpdateUser(ctx, tx, user); err != nil {
		return nil, exception.ToErrorMsg(err.Error(), err)
	} else {
		return response.ToUserResponse(updateUser), nil
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
		userResponses := response.ToUserResponses(user)
		return userResponses, nil
	}
}

func (u *UserServiceImpl) ChangePassword(ctx context.Context, req *request.ChangePasswordRequest) *exception.ErrorMsg {

	tx, err := u.DB.Begin()
	if err != nil {
		return exception.ToErrorMsg(err.Error(), err)
	}
	defer helpers.RollbackCommit(tx)

	if findById, err := u.UserRepository.FindByIdUser(ctx, u.DB, req.ID); err != nil {
		return exception.ToErrorMsg("Data User not found", err)
	} else {
		if findById.Password != req.PreviousPassword {
			return exception.ToErrorMsg("The current password does not match.", exception.BadRequest)
		}
	}

	if req.Password != req.ConfirmPassword {
		return exception.ToErrorMsg("Confirm Password not match with New Password", exception.BadRequest)
	}

	user := &domain.Users{
		ID:       req.ID,
		Password: req.Password,
	}

	if isSuccess, err := u.UserRepository.ChangePassword(ctx, tx, user); err != nil {
		return exception.ToErrorMsg(err.Error(), err)
	} else {
		if isSuccess {
			return nil
		} else {
			return exception.ToErrorMsg(err.Error(), err)
		}
	}
}

func (u *UserServiceImpl) FindByUsername(ctx context.Context, username string) (bool, *exception.ErrorMsg) {
	findByUsername, err := u.UserRepository.FindByUsername(ctx, u.DB, username)
	if findByUsername == nil {
		return false, exception.ToErrorMsg(err.Error(), err)
	}
	return true, nil
}
