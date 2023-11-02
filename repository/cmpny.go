package repository

import (
	"errors"
	"graph/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateCompany(cmpnydetails models.Company) (models.Company, error) {
	result := r.DB.Create(&cmpnydetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Company{}, errors.New("not creating")
	}
	return cmpnydetails, nil
}

func (r *Repo) ViewCompanies() ([]models.Company, error) {
	var cmpy []models.Company
	result := r.DB.Find(&cmpy)
	if result.Error != nil {
		return nil, errors.New("data not there")
	}
	return cmpy, nil
}

func (r *Repo) ViewCompanyById(Id string) (models.Company, error) {
	var cmpy models.Company
	result := r.DB.Where("id=?", Id).Find(&cmpy)
	if result.Error != nil {
		return models.Company{}, errors.New("companyid not found")
	}
	return cmpy, nil
}
func (r *Repo) CreateJob(jbdetails models.Job) (models.Job, error) {
	result := r.DB.Create(&jbdetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.Job{}, errors.New("not creating")
	}
	return jbdetails, nil
}
func (r *Repo) ViewAllJobs() ([]models.Job, error) {
	var jb []models.Job
	result := r.DB.Find(&jb)
	if result.Error != nil {
		return nil, errors.New("data not found")
	}
	return jb, nil
}
func (r *Repo) ViewJobByCid(Id string) ([]*models.Job, error) {
	var jobDetails []*models.Job
	result := r.DB.Where("cid=?", Id).Find(&jobDetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return nil, errors.New("could not find job")
	}
	return jobDetails, nil
}
