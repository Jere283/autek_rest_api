package service

import (
	"github.com/Jere283/autek_rest_api/models"
	"github.com/Jere283/autek_rest_api/repository"
)

type UserService struct {
	UserRepo repository.UserRepository
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	return s.UserRepo.FindById(id)
}
