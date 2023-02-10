package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
)

type ExpeditionRepository interface {
	AddExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (bool, error)
	UpdateExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (isSuccess bool, isFailed error)
	DeleteExpedition(ctx context.Context, tx *sql.Tx, expID int) (isSuccess bool, isFailed error)
	FindAllExpedition(ctx context.Context, db *sql.DB) (expeditions []*domain.Expedition, isFailed error)
	FindExpeditionByID(ctx context.Context, db *sql.DB, expID int) (expeditions *domain.Expedition, isFailed error)
}

type ExpeditionRepositoryImpl struct {
}

func NewExpeditionRepositoryImpl() ExpeditionRepository {
	return &ExpeditionRepositoryImpl{}
}

func (e *ExpeditionRepositoryImpl) AddExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (bool, error) {
	if result, err := tx.ExecContext(ctx, "INSERT INTO expedition (name) VALUES(?)", &expedition.Name); err != nil {
		return false, err
	} else {
		ID, _ := result.LastInsertId()
		expedition.ID = int(ID)
		return true, nil
	}
}

func (e *ExpeditionRepositoryImpl) UpdateExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (bool, error) {
	if _, err := tx.ExecContext(ctx, "UPDATE expedition SET name = ? WHERE id = ?", expedition.Name, expedition.ID); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (e *ExpeditionRepositoryImpl) DeleteExpedition(ctx context.Context, tx *sql.Tx, expID int) (isSuccess bool, isFailed error) {
	if _, err := tx.ExecContext(ctx, "DELETE FROM expedition WHERE id = ?", expID); err != nil {
		return false, err
	} else {
		return true, nil
	}
}

func (e *ExpeditionRepositoryImpl) FindAllExpedition(ctx context.Context, db *sql.DB) ([]*domain.Expedition, error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, name FROM expedition ORDER BY id"); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		var expeditions []*domain.Expedition
		for rows.Next() {
			var expedition domain.Expedition
			if err := rows.Scan(&expedition.ID, &expedition.Name); err != nil {
				return nil, err
			}
			expeditions = append(expeditions, &expedition)
		}
		return expeditions, nil
	}
}

func (e *ExpeditionRepositoryImpl) FindExpeditionByID(ctx context.Context, db *sql.DB, expID int) (expeditions *domain.Expedition, isFailed error) {
	if rows, err := db.QueryContext(ctx, "SELECT id, name FROM expedition WHERE id = ?", expID); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		var expedition domain.Expedition
		if rows.Next() {

			if err := rows.Scan(&expedition.ID, &expedition.Name); err != nil {
				return nil, errors.New(fmt.Sprintf("Data with ID-%d not found.", expedition.ID))
			} else {
				fmt.Println(expedition)
				return &expedition, nil
			}
		} else {
			return nil, err
		}
	}
}
