package repository

import (
	"github.com/Jere283/autek_rest_api/database"
	"github.com/Jere283/autek_rest_api/models"
)

type UserRepository struct{}

func (r *UserRepository) FindById(id string) (models.User, error) {
	var user models.User
	err := database.DB.First(&user, "id_user=?", id).Error
	return user, err
}
