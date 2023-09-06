package service

import (
	"vnet-mysql/pkg/models"
	"vnet-mysql/pkg/repository"
)

type UserService struct {
	UserRepository *repository.UserRepository
}

func NewUserService(ur *repository.UserRepository) *UserService {
	return &UserService{UserRepository: ur}
}

func (us *UserService) CreateUser(user *models.User) error {
	return us.UserRepository.Create(user)
}

func (us *UserService) GetUserByID(id uint) (*models.User, error) {
	return us.UserRepository.FindByID(id)
}

func (us *UserService) UpdateUser(user *models.User) error {
	return us.UserRepository.Update(user)
}

func (us *UserService) DeleteUser(user *models.User) error {
	return us.UserRepository.Delete(user)
}

func (us *UserService) GetAllUsers() ([]models.User, error) {
	return us.UserRepository.FindAll()
}
