package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
	"github.com/dimassfeb-09/restapi-ecommerce.git/exception"
)

type CityRepository interface {
	CreateCity(ctx context.Context, db *sql.Tx, cityName string, provinceId int) (*domain.City, error)
	UpdateCity(ctx context.Context, db *sql.Tx, city *domain.City) (*domain.City, error)
	DeleteCity(ctx context.Context, db *sql.Tx, cityId int) error
	FindCityById(ctx context.Context, db *sql.DB, cityId int) (*domain.City, error)
	FindCityByName(ctx context.Context, db *sql.DB, cityName string) (*domain.City, error)
	FindCityByProvinceID(ctx context.Context, db *sql.DB, provinceId int) (bool, error)
	FindAllCity(ctx context.Context, db *sql.DB) ([]*domain.City, error)
}

type CityRepositoryImpl struct {
}

func NewCityRepositoryImpl() CityRepository {
	return &CityRepositoryImpl{}
}

func (c *CityRepositoryImpl) CreateCity(ctx context.Context, db *sql.Tx, cityName string, provinceId int) (*domain.City, error) {
	if result, err := db.ExecContext(ctx, "INSERT INTO city(name, province_id) VALUES (?, ?)", cityName, provinceId); err != nil {
		return nil, err
	} else {
		if id, err := result.LastInsertId(); err != nil {
			return nil, err
		} else {
			city := &domain.City{
				ID:         int(id),
				Name:       cityName,
				ProvinceID: provinceId,
			}
			return city, nil
		}
	}
}

func (c *CityRepositoryImpl) UpdateCity(ctx context.Context, db *sql.Tx, city *domain.City) (*domain.City, error) {
	if _, err := db.ExecContext(ctx, "UPDATE city SET name = ?, province_id = ? WHERE id = ?", &city.Name, &city.ProvinceID, &city.ID); err != nil {
		return nil, err
	} else {
		return city, nil
	}
}

func (c *CityRepositoryImpl) DeleteCity(ctx context.Context, db *sql.Tx, cityId int) error {
	if _, err := db.ExecContext(ctx, "DELETE FROM city WHERE id = ?", cityId); err != nil {
		return err
	} else {
		return nil
	}
}

func (c *CityRepositoryImpl) FindCityById(ctx context.Context, db *sql.DB, cityId int) (*domain.City, error) {
	if rows, err := db.QueryContext(ctx, "SELECT * FROM city WHERE id = ?", cityId); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var city domain.City
		if rows.Next() {
			if err := rows.Scan(&city.ID, &city.Name, &city.ProvinceID); err != nil {
				return nil, err
			} else {
				return &city, err
			}
		} else {
			return nil, exception.ErrorNotFound // has error
		}
	}
}

func (c *CityRepositoryImpl) FindCityByName(ctx context.Context, db *sql.DB, cityName string) (*domain.City, error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, name FROM city WHERE name = ?", cityName); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var city domain.City
		if rows.Next() {
			if err := rows.Scan(&city.ID, &city.Name); err != nil {
				return nil, err
			}
			return &city, nil
		} else {
			return nil, err
		}
	}
}

func (c *CityRepositoryImpl) FindCityByProvinceID(ctx context.Context, db *sql.DB, provinceId int) (bool, error) {

	if rows, err := db.QueryContext(ctx, "SELECT id FROM city WHERE province_id = ?", provinceId); err != nil {
		return false, err
	} else {
		defer rows.Close()

		var city domain.City
		if rows.Next() {
			if err := rows.Scan(&city.ID); err != nil {
				return false, err
			}
			return true, nil
		} else {
			return false, err
		}
	}
}

func (c *CityRepositoryImpl) FindAllCity(ctx context.Context, db *sql.DB) ([]*domain.City, error) {

	if rows, err := db.QueryContext(ctx, "SELECT  * FROM city ORDER BY id ASC"); err != nil {
		return nil, err
	} else {
		defer rows.Close()

		var cities []*domain.City
		for rows.Next() {
			var city domain.City
			if err := rows.Scan(&city.ID, &city.Name, &city.ProvinceID); err != nil {
				return nil, err
			} else {
				cities = append(cities, &city)
			}
		}
		return cities, nil
	}
}
