package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
)

type ProvinceRepository interface {
	CreateProvince(ctx context.Context, tx *sql.Tx, name string) (*domain.Province, error)
	UpdateProvince(ctx context.Context, tx *sql.Tx, province *domain.Province) (*domain.Province, error)
	DeleteProvince(ctx context.Context, tx *sql.Tx, provinceId int) error
	FindProvinceByName(ctx context.Context, db *sql.DB, name string) (*domain.Province, error)
	FindProvinceById(ctx context.Context, db *sql.DB, provinceId int) (*domain.Province, error)
	FindAllProvince(ctx context.Context, db *sql.DB) ([]*domain.Province, error)
}

type ProvinceRepositoryImpl struct {
}

func NewProvinceRepositoryImpl() ProvinceRepository {
	return &ProvinceRepositoryImpl{}
}

func (p *ProvinceRepositoryImpl) CreateProvince(ctx context.Context, tx *sql.Tx, name string) (*domain.Province, error) {
	if result, err := tx.ExecContext(ctx, "INSERT INTO province(name) VALUES (?)", name); err != nil {
		return nil, err
	} else {
		id, _ := result.LastInsertId()
		province := &domain.Province{
			ID:   int(id),
			Name: name,
		}
		return province, nil
	}
}

func (p *ProvinceRepositoryImpl) UpdateProvince(ctx context.Context, tx *sql.Tx, province *domain.Province) (*domain.Province, error) {
	tx.ExecContext(ctx, "UPDATE province SET name = ? WHERE id = ?;", &province.Name, &province.ID)
	return province, nil
}

func (p *ProvinceRepositoryImpl) DeleteProvince(ctx context.Context, tx *sql.Tx, provinceId int) error {
	if _, err := tx.ExecContext(ctx, "DELETE FROM province WHERE id = ?", provinceId); err != nil {
		return err
	} else {
		return nil
	}
}

func (p *ProvinceRepositoryImpl) FindProvinceById(ctx context.Context, db *sql.DB, provinceId int) (*domain.Province, error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, name FROM province WHERE id = ?", provinceId); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var province domain.Province
		if rows.Next() {
			if err := rows.Scan(&province.ID, &province.Name); err != nil {
				return nil, err
			} else {
				return &province, err
			}
		} else {
			return nil, exception.ErrorNotFound // has error
		}
	}
}

func (p *ProvinceRepositoryImpl) FindProvinceByName(ctx context.Context, db *sql.DB, name string) (*domain.Province, error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, name FROM province WHERE name = ?", name); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var province domain.Province
		if rows.Next() {
			if err := rows.Scan(&province.ID, &province.Name); err != nil {
				return nil, err
			}
			return &province, nil
		} else {
			return nil, err
		}
	}
}

func (p *ProvinceRepositoryImpl) FindAllProvince(ctx context.Context, db *sql.DB) ([]*domain.Province, error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, name FROM province ORDER BY id"); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var provinces []*domain.Province
		for rows.Next() {
			var province domain.Province
			if err := rows.Scan(&province.ID, &province.Name); err != nil {
				return nil, err
			} else {
				provinces = append(provinces, &province)
			}
		}

		return provinces, nil
	}
}
