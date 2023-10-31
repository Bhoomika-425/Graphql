package service

import (
	"graph/database"
	"graph/graph/model"
	newModel "graph/models"
	"strconv"
)

func (s *Service) UserSignup(userData model.NewUser) (*model.User, error) {
	hashedPassword, err := database.HashPassword(userData.Password)
	if err != nil {
		return nil, err
	}

	userDetails := newModel.User{
		Username:   userData.Username,
		Email:      userData.Email,
		Hashedpass: hashedPassword,
	}

	userDetails, err = s.userRepo.CreateSignup(userDetails)
	if err != nil {
		return nil, err
	}

	uid := strconv.FormatUint(uint64(userDetails.ID), 10)

	return &model.User{
		ID:        uid,
		Username:  userDetails.Username,
		Email:     userDetails.Email,
		CreatedAt: userDetails.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: userDetails.UpdatedAt.Format("2006-01-02 15:04:05"),
	}, nil
}
