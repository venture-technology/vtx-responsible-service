package repository

import (
	"context"
	"database/sql"

	"github.com/venture-technology/vtx-responsible-service/models"
)

type IChildRepository interface {
	CreateChild(ctx context.Context, child *models.Child) error
	GetChild(ctx context.Context, rg *string) (*models.Child, error)
	FindAllChildren(ctx context.Context, cpf *string) ([]models.Child, error)
	UpdateChild(ctx context.Context, rg *string) error
	DeleteChild(ctx context.Context, rg *string) error
}

type ChildRepository struct {
	db *sql.DB
}

func NewChildRepository(conn *sql.DB) *ChildRepository {
	return &ChildRepository{
		db: conn,
	}
}

func (cr *ChildRepository) CreateChild(ctx context.Context, child *models.Child) error {
	return nil
}

func (cr *ChildRepository) GetChild(ctx context.Context, rg *string) (*models.Child, error) {
	return nil, nil
}

func (cr *ChildRepository) FindAllChildren(ctx context.Context, cpf *string) ([]models.Child, error) {
	return nil, nil
}

func (cr *ChildRepository) UpdateChild(ctx context.Context, rg *string) error {
	return nil
}

func (cr *ChildRepository) DeleteChild(ctx context.Context, rg *string) error {
	return nil
}
