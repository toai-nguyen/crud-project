package repositories

import (
	"example.com/go-back-end/config"
	"example.com/go-back-end/models"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) Create(user models.User) (models.User, error) {
	result := config.DB.Create(&user)

	return user, result.Error
}
func (r *UserRepository) Update(user models.User) (models.User, error) {
	result := config.DB.Save(&user)

	return user, result.Error
}

func (r *UserRepository) FindAll() ([]models.User, error) {
	var users []models.User
	result := config.DB.Find(&users)

	return users, result.Error
}

func (r *UserRepository) FindById(id string) (models.User, error) {
	var user models.User
	result := config.DB.First(&user, id)

	return user, result.Error
}

func (r *UserRepository) Delete(user models.User) error {
	result := config.DB.Delete(&user)

	return result.Error
}
