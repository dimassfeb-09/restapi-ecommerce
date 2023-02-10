package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
)

type AuthRepository interface {
	AuthLogin(ctx context.Context, db *sql.DB, user *domain.AuthUser) (*domain.AuthUser, error)
	AuthRegister(ctx context.Context, db *sql.Tx, user *domain.AuthUser) (bool, error)
}

type AuthRepositoryImpl struct {
}

func NewAuthRepositoryImpl() AuthRepository {
	return &AuthRepositoryImpl{}
}

func (a *AuthRepositoryImpl) AuthLogin(ctx context.Context, db *sql.DB, user *domain.AuthUser) (*domain.AuthUser, error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, username, password FROM users WHERE username = ? AND password = ?", user.Username, user.Password); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var authUser domain.AuthUser
		if rows.Next() {
			if err := rows.Scan(&authUser.ID, &authUser.Username, &authUser.Password); err != nil {
				return nil, err
			}
			return &authUser, nil
		} else {
			return nil, errors.New("Username or Password wrong.")
		}
	}
}

func (a *AuthRepositoryImpl) AuthRegister(ctx context.Context, tx *sql.Tx, user *domain.AuthUser) (bool, error) {
	fmt.Println("gak ke eksekusi")
	sqlQuery := "INSERT INTO users(name, username, password) VALUES(?, ?, ?)"
	if _, err := tx.ExecContext(ctx, sqlQuery, user.Name, user.Username, user.Password); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
