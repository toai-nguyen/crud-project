package services

import (
	"example.com/go-back-end/dto/request"
	"example.com/go-back-end/models"
	"example.com/go-back-end/repositories"
)

type UserService struct {
	userRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	return s.userRepo.FindAll()
}

func (s *UserService) GetUserByID(id string) (models.User, error) {
	return s.userRepo.FindById(id)
}

func (s *UserService) CreateUser(req request.CreateUserRequest) (models.User, error) {
	user := models.User{
		Name:     req.Name,
		Email:    &req.Email,
		Phone:    &req.Phone,
		Password: req.Password,
		Role:     "user",
	}
	return s.userRepo.Create(user)
}

func (s *UserService) UpdateUser(id string, req request.UpdateUserRequest) (models.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return user, err
	}

	user.Name = req.Name
	if req.Email != "" {
		*user.Email = req.Email
	}

	return s.userRepo.Update(user)
}

func (s *UserService) DeleteUser(id string) error {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return err
	}

	return s.userRepo.Delete(user)
}
