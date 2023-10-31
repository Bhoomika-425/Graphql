package service

import (
	"graph/graph/model"
	newModel "graph/models"
	"strconv"
)

func (s *Service) UserCompany(cmdata model.NewCompany) (*model.Company, error) {
	companydetails := newModel.Company{
		Name:     cmdata.Name,
		Location: cmdata.Loc,
		Salary:   cmdata.Sal,
	}
	companydetails, err := s.userRepo.CreateCompany(companydetails)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		ID:   strconv.FormatUint(uint64(companydetails.ID), 10),
		Name: cmdata.Name,
		Loc:  cmdata.Loc,
		Sal:  cmdata.Sal,
	}, nil

}

func (s *Service) GetCompanies() ([]*model.Company, error) {
	cmp, err := s.userRepo.ViewCompanies()
	if err != nil {
		return nil, err
	}

	cds := []*model.Company{}

	for _, v := range cmp {
		cd := model.Company{
			ID:   strconv.FormatUint(uint64(v.ID), 10),
			Name: v.Name,
			Loc:  v.Location,
			Sal:  v.Salary,
		}
		cds = append(cds, &cd)
	}
	return cds, nil
}

func (s *Service) GetCompanyById(Id string) (*model.Company, error) {
	cmpnydetails, err := s.userRepo.ViewCompanyById(Id)
	if err != nil {
		return nil, err
	}
	return &model.Company{
		ID:   strconv.FormatUint(uint64(cmpnydetails.ID), 10),
		Name: cmpnydetails.Name,
		Loc:  cmpnydetails.Location,
		Sal:  cmpnydetails.Salary,
	}, nil
}
