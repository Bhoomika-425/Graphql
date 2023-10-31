package repository

import (
	"errors"
	"graph/models"

	"gorm.io/gorm"
)

type Repo struct {
	DB *gorm.DB
}

type UserRepo interface {
	CreateSignup(input models.User) (models.User, error)
	CreateCompany(input models.Company) (models.Company, error)
	ViewCompanies() ([]models.Company, error)
	ViewCompanyById(Id string) (models.Company, error)
}

func NewRepo(db *gorm.DB) (UserRepo, error) {
	if db == nil {
		return nil, errors.New("db is null")
	}
	return &Repo{
		DB: db,
	}, nil
}
