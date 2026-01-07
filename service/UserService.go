package service

import (
	"ManageTask/models"
	"ManageTask/repository"
	"ManageTask/utils"
	"errors"
	"strings"
)

type UserService struct {
	userRepo *repository.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		userRepo: repository.NewUserRepository(repository.DB),
	}
}

//func (s *UserService) CheckUser(email string) (bool, error) {
//	exists, err := s.userRepo.ExitByEmail(email)
//	if err != nil {
//		return false, err
//	}
//	return exists, nil
//}

func (s *UserService) Register(email, password string) error {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" || password == "" {
		return errors.New("email/password cannot be empty")
	}

	exit, err := s.userRepo.ExitByEmail(email)
	if err != nil {
		return err
	}
	if exit {
		return errors.New("email already exit")
	}

	hash, err := utils.Hash(password)
	if err != nil {
		return err
	}

	user := &models.User{
		Email:        email,
		PasswordHash: hash,
		Role:         "user",
	}

	return s.userRepo.Create(user)
}

func (s *UserService) Login(email, password string) (*models.User, error) {
	email = strings.TrimSpace(strings.ToLower(email))
	if email == "" || password == "" {
		return nil, errors.New("invalid email or password")
	}

	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	if err := utils.Check(user.PasswordHash, password); err != nil {
		return nil, errors.New("invalid email or password")
	}

	return user, nil
}
