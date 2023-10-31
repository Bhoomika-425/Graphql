package repository

import (
	"errors"

	"graph/models"

	"github.com/rs/zerolog/log"
)

func (r *Repo) CreateSignup(userdetails models.User) (models.User, error) {
	result := r.DB.Create(&userdetails)
	if result.Error != nil {
		log.Info().Err(result.Error).Send()
		return models.User{}, errors.New("not creating")
	}
	return userdetails, nil
}
