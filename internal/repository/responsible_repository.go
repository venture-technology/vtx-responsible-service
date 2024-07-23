package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/venture-technology/vtx-responsible-service/models"
	"github.com/venture-technology/vtx-responsible-service/utils"
)

type IResponsibleRepository interface {
	CreateResponsible(ctx context.Context, responsible *models.Responsible) error
	GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error)
	UpdateResponsible(ctx context.Context, responsible *models.Responsible) error
	DeleteResponsible(ctx context.Context, cpf *string) error
	AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error)
}

type ResponsibleRepository struct {
	db *sql.DB
}

func NewResponsibleRepository(conn *sql.DB) *ResponsibleRepository {
	return &ResponsibleRepository{
		db: conn,
	}
}

func (rer *ResponsibleRepository) CreateResponsible(ctx context.Context, responsible *models.Responsible) error {
	sqlQuery := `INSERT INTO responsible (name, email, password, cpf, street, number, zip, complement, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := rer.db.Exec(sqlQuery, responsible.Name, responsible.Email, responsible.Password, responsible.CPF, responsible.Street, responsible.Number, responsible.ZIP, responsible.Complement, responsible.Status)
	return err
}

func (rer *ResponsibleRepository) GetResponsible(ctx context.Context, cpf *string) (*models.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, email, street, number, zip, status, complement FROM responsible WHERE cpf = $1 LIMIT 1`
	var responsible models.Responsible
	err := rer.db.QueryRow(sqlQuery, *cpf).Scan(
		&responsible.ID,
		&responsible.Name,
		&responsible.CPF,
		&responsible.Email,
		&responsible.Street,
		&responsible.Number,
		&responsible.ZIP,
		&responsible.Status,
		&responsible.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	return &responsible, nil
}

func (rer *ResponsibleRepository) UpdateResponsible(ctx context.Context, responsible *models.Responsible) error {
	sqlQuery := `SELECT name, email, password, street, number, zip, status, complement FROM responsible WHERE cpf = $1 LIMIT 1`
	var currentResponsible models.Responsible
	err := rer.db.QueryRow(sqlQuery, responsible.CPF).Scan(
		&currentResponsible.Name,
		&currentResponsible.Email,
		&currentResponsible.Password,
		&currentResponsible.Street,
		&currentResponsible.Number,
		&currentResponsible.ZIP,
		&currentResponsible.Status,
		&currentResponsible.Complement,
	)
	if err != nil || err == sql.ErrNoRows {
		return err
	}

	if responsible.Name != "" && responsible.Name != currentResponsible.Name {
		currentResponsible.Name = responsible.Name
	}

	if responsible.Email != "" && responsible.Email != currentResponsible.Email {
		currentResponsible.Email = responsible.Email
	}
	if responsible.Password != "" && responsible.Password != currentResponsible.Password {
		currentResponsible.Password = responsible.Password
		currentResponsible.Password = utils.HashPassword(currentResponsible.Password)
	}
	if responsible.Street != "" && responsible.Street != currentResponsible.Street {
		currentResponsible.Street = responsible.Street
	}
	if responsible.Number != "" && responsible.Number != currentResponsible.Number {
		currentResponsible.Number = responsible.Number
	}
	if responsible.ZIP != "" && responsible.ZIP != currentResponsible.ZIP {
		currentResponsible.ZIP = responsible.ZIP
	}
	if responsible.Complement != "" && responsible.Complement != currentResponsible.Complement {
		currentResponsible.Complement = responsible.Complement
	}

	sqlQueryUpdate := `UPDATE responsible SET name = $1, email = $2, password = $3, street = $4, number = $5, zip = $6, complement = $7 WHERE cpf = $8`
	_, err = rer.db.ExecContext(ctx, sqlQueryUpdate, currentResponsible.Name, currentResponsible.Email, currentResponsible.Password, currentResponsible.Street, currentResponsible.Number, currentResponsible.ZIP, currentResponsible.Complement, responsible.CPF)
	return err
}

func (rer *ResponsibleRepository) DeleteResponsible(ctx context.Context, cpf *string) error {
	tx, err := rer.db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		} else if err != nil {
			_ = tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()
	_, err = tx.Exec("DELETE FROM responsible WHERE cpf = $1", *cpf)
	return err
}

func (rer *ResponsibleRepository) AuthResponsible(ctx context.Context, responsible *models.Responsible) (*models.Responsible, error) {
	sqlQuery := `SELECT id, name, cpf, street, email, number, zip, status, password FROM responsible WHERE email = $1 LIMIT 1`
	var responsibleData models.Responsible
	err := rer.db.QueryRow(sqlQuery, responsible.Email).Scan(
		&responsibleData.ID,
		&responsibleData.Name,
		&responsibleData.CPF,
		&responsibleData.Street,
		&responsibleData.Email,
		&responsibleData.Number,
		&responsibleData.ZIP,
		&responsibleData.Status,
		&responsibleData.Password,
	)
	if err != nil || err == sql.ErrNoRows {
		return nil, err
	}
	match := responsibleData.Password == responsible.Password
	if !match {
		return nil, fmt.Errorf("email or password wrong")
	}
	responsibleData.Password = ""
	return &responsibleData, nil
}
