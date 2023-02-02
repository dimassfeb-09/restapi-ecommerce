package repository

import (
	"context"
	"database/sql"
	"github.com/dimassfeb-09/restapi-ecommerce.git/entity/domain"
)

type ExpeditionRepository interface {
	AddExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (*domain.Expedition, error)
	UpdateExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (isSuccess bool, isFailed error)
	DeleteExpedition(ctx context.Context, tx *sql.Tx, expID int) (isSuccess bool, isFailed error)
	FindAllExpedition(ctx context.Context, tx *sql.Tx) (expeditions []*domain.Expedition, isFailed error)
	FindExpeditionByID(ctx context.Context, tx *sql.Tx, expID int) (expeditions *domain.Expedition, isFailed error)
}

type ExpeditionRepositoryImpl struct {
}

func NewExpeditionRepositoryImpl() ExpeditionRepository {
	return &ExpeditionRepositoryImpl{}
}

func (e *ExpeditionRepositoryImpl) AddExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (*domain.Expedition, error) {
	if result, err := tx.ExecContext(ctx, "INSERT INTO expedition (name) VALUES(?)", &expedition.Name); err != nil {
		return nil, err
	} else {
		ID, _ := result.LastInsertId()
		expedition.ID = int(ID)
		return expedition, nil
	}
}

func (e *ExpeditionRepositoryImpl) UpdateExpedition(ctx context.Context, tx *sql.Tx, expedition *domain.Expedition) (bool, error) {
	if _, err := tx.ExecContext(ctx, "UPDATE expedition SET name = ? WHERE id = ?", &expedition.Name, &expedition.ID); err != nil {
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

func (e *ExpeditionRepositoryImpl) FindAllExpedition(ctx context.Context, tx *sql.Tx) ([]*domain.Expedition, error) {
	if rows, err := tx.QueryContext(ctx, "SELECT (id, name) FROM expedition ORDER BY id"); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		var expeditions []*domain.Expedition
		for rows.Next() {
			var expedition *domain.Expedition
			if err := rows.Scan(&expedition.ID, &expedition.Name); err != nil {
				return nil, err
			}
			expeditions = append(expeditions, expedition)
		}
		return expeditions, nil
	}
}

func (e *ExpeditionRepositoryImpl) FindExpeditionByID(ctx context.Context, tx *sql.Tx, expID int) (expeditions *domain.Expedition, isFailed error) {
	if rows, err := tx.QueryContext(ctx, "SELECT (id, name) FROM expedition WHERE id = ?", expID); err != nil {
		return nil, err
	} else {
		defer rows.Close()
		var expedition *domain.Expedition
		if rows.Next() {
			if err := rows.Scan(&expedition.ID, &expedition.Name); err != nil {
				return nil, err
			}
		}
		return expedition, nil
	}
}
