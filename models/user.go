package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username   string `json:"username" validate:"required"`
	Email      string `json:"email" validate:"required"`
	Hashedpass string `json:"hashed_pass" validate:"required"`
}
type Company struct {
	gorm.Model
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Salary   string `json:"salary" validate:"required"`
}
type Job struct {
	gorm.Model
	Company      Company `gorm:"foreignKey:Cid"`
	Cid          string  `json:"cid"`
	Jobname      string  `json:"name"`
	Salary       string  `json:"salary"`
	NoticePeriod string  `json:"notice_period"`
}
