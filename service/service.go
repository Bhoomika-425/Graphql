package service

import (
	"errors"
	"graph/graph/model"
	"graph/repository"
)

type UserService interface {
	UserSignup(model.NewUser) (*model.User, error)
	UserCompany(model.NewCompany)(*model.Company,error)
	GetCompanies()([]*model.Company,error)
	GetCompanyById(Id string)(*model.Company,error)
}
type Service struct {
	userRepo repository.UserRepo
}

func NewService(userRepo repository.UserRepo) (UserService, error) {
	if userRepo == nil {
		return nil, errors.New("interface cannot be null")
	}
	return &Service{
		userRepo: userRepo,
	}, nil
}

