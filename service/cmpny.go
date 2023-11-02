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
func (s *Service) UserJob(jbdata model.NewJob) (*model.Job, error) {
	jbdetails := newModel.Job{
		Cid:          jbdata.Cid,
		Jobname:      jbdata.Jobname,
		Salary:       jbdata.Salary,
		NoticePeriod: jbdata.Noticeperiod,
	}
	jbdetails, err := s.userRepo.CreateJob(jbdetails)
	if err != nil {
		return nil, err
	}
	return &model.Job{
		ID:           strconv.FormatUint(uint64(jbdetails.ID), 10),
		Cid:          jbdata.Cid,
		Jobname:      jbdata.Jobname,
		Salary:       jbdata.Salary,
		Noticeperiod: jbdata.Noticeperiod,
	}, nil
}

func (s *Service) GetAllJobs() ([]*model.Job, error) {
	jb, err := s.userRepo.ViewAllJobs()
	if err != nil {
		return nil, err
	}
	var jbs []*model.Job

	for _, a := range jb {
		jobData := &model.Job{
			ID:           strconv.FormatUint(uint64(a.ID), 10),
			Cid:          a.Cid,
			Jobname:      a.Jobname,
			Salary:       a.Salary,
			Noticeperiod: a.NoticePeriod,
		}
		jbs = append(jbs, jobData)
	}
	return jbs, nil

}
func (s *Service) GetJobByCid(cid string) ([]*model.Job, error) {
	job, err := s.userRepo.ViewJobByCid(cid)
	if err != nil {
		return nil, err
	}
	var jobDatas []*model.Job
	for _, v := range job {
		jobData := &model.Job{
			ID:           strconv.FormatUint(uint64(v.ID), 10),
			Cid:          v.Cid,
			Jobname:      v.Jobname,
			Salary:       v.Salary,
			Noticeperiod: v.NoticePeriod,
		}
		jobDatas = append(jobDatas, jobData)
	}
	return jobDatas, nil

}

// return &model.Job{
// 	ID:           strconv.FormatUint(uint64(jbdetails.ID), 10),
// 	Jobname:      jbdetails.Jobname,
// 	Salary:       jbdetails.Salary,
// 	Noticeperiod: jbdetails.NoticePeriod,
// }, nil
