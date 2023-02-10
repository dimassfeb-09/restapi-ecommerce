package repository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
)

type UserRepository interface {
	CreateUser(ctx context.Context, DB *sql.Tx, users *domain.Users) (*domain.Users, error)
	FindByIdUser(ctx context.Context, DB *sql.DB, userId int) (*domain.Users, error)
	FindByUsername(ctx context.Context, DB *sql.DB, username string) (*domain.Users, error)
	UpdateUser(ctx context.Context, DB *sql.Tx, users *domain.Users) (*domain.Users, error)
	DeleteUser(ctx context.Context, DB *sql.Tx, userId int) error
	FindAll(ctx context.Context, DB *sql.DB) ([]*domain.Users, error)
	ChangePassword(ctx context.Context, db *sql.Tx, users *domain.Users) (bool, error)
}

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() UserRepository {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) CreateUser(ctx context.Context, DB *sql.Tx, users *domain.Users) (*domain.Users, error) {

	var Sql string = "INSERT INTO users(name, username, password) VALUES (?, ?, ?)"
	rows, err := DB.ExecContext(ctx, Sql, users.Name, users.Username, users.Password)
	if err != nil {
		errMsg := fmt.Errorf("Error Create User: %s", err.Error())
		return nil, errMsg
	}

	lastId, err := rows.LastInsertId() // Getting last Id in Database
	if err != nil {
		errMsg := fmt.Errorf("Error Last ID: %s", err.Error())
		return nil, errMsg
	}

	user := &domain.Users{
		ID:        int(lastId),
		Name:      users.Name,
		Username:  users.Username,
		Balance:   0,
		CreatedAt: users.CreatedAt,
	}

	return user, nil
}

func (u *UserRepositoryImpl) FindByIdUser(ctx context.Context, DB *sql.DB, userId int) (*domain.Users, error) {
	rows, err := DB.QueryContext(ctx, "SELECT * FROM users WHERE id = ?", userId)
	if err != nil {
		errMsg := fmt.Errorf("Error Query Find User By ID: %s", err.Error())
		return nil, errMsg
	}
	defer rows.Close()

	var user domain.Users
	if rows.Next() { // if data founded
		err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Balance, &user.CreatedAt)
		if err != nil {
			errMsg := fmt.Errorf("Error Scan User Find By ID: %s", err)
			return nil, errMsg
		}
		return &user, nil
	} else { // if data not founded
		errMsg := fmt.Errorf("User dengan ID %d tidak ditemukan.", userId)
		return nil, errMsg
	}
}

func (u *UserRepositoryImpl) FindByUsername(ctx context.Context, DB *sql.DB, username string) (*domain.Users, error) {
	rows, err := DB.QueryContext(ctx, "SELECT id, username FROM users WHERE username = ?", username)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var user domain.Users
	if rows.Next() {
		if err := rows.Scan(&user.ID, &user.Username); err != nil {
			return nil, err
		}
		return &user, nil
	} else {
		errMsg := fmt.Errorf("Username %s tidak ditemukan.", username)
		return nil, errMsg
	}

}

func (u *UserRepositoryImpl) UpdateUser(ctx context.Context, DB *sql.Tx, users *domain.Users) (*domain.Users, error) {
	SQL := "UPDATE users SET name = ?, username = ?, password = ?, balance = ? WHERE id = ?"

	_, err := DB.ExecContext(ctx, SQL, &users.Name, &users.Username, &users.Password, &users.Balance, &users.ID)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) DeleteUser(ctx context.Context, DB *sql.Tx, userId int) error {
	if _, err := DB.ExecContext(ctx, "DELETE FROM users WHERE id = ?", userId); err != nil {
		return err
	} else {
		return nil
	}
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, DB *sql.DB) ([]*domain.Users, error) {
	rows, err := DB.QueryContext(ctx, "SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*domain.Users
	for rows.Next() {
		var user domain.Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Username, &user.Password, &user.Balance, &user.CreatedAt); err != nil {
			return nil, err
		} else {
			users = append(users, &user)
		}
	}

	return users, nil
}

func (u *UserRepositoryImpl) ChangePassword(ctx context.Context, db *sql.Tx, users *domain.Users) (bool, error) {
	if _, err := db.ExecContext(ctx, "UPDATE users SET password = ? WHERE id = ?", &users.Password, &users.ID); err != nil {
		return false, err
	} else {
		return true, nil
	}
}
