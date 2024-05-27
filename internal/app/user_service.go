package app

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"log"
)

type UserService interface {
	FindByEmail(email string) (domain.User, error)
	FindById(id uint64) (domain.User, error)
	Find(id uint64) (interface{}, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint64) error
}

type userService struct {
	userRepo database.UserRepository
}

func NewUserService(ur database.UserRepository) UserService {
	return userService{
		userRepo: ur,
	}
}

func (s userService) FindByEmail(email string) (domain.User, error) {
	user, err := s.userRepo.FindByEmail(email)
	if err != nil {
		log.Printf("UserService: %s", err)
		return domain.User{}, err
	}

	return user, err
}

func (s userService) FindById(id uint64) (domain.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		log.Printf("UserService: %s", err)
		return domain.User{}, err
	}
	return user, err
}

func (s userService) Find(id uint64) (interface{}, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		log.Printf("UserService: %s", err)
		return domain.User{}, err
	}
	return user, err
}

func (s userService) Update(user domain.User) (domain.User, error) {
	user, err := s.userRepo.Update(user)
	if err != nil {
		log.Printf("UserService: %s", err)
		return domain.User{}, err
	}

	return user, nil
}

func (s userService) Delete(id uint64) error {
	err := s.userRepo.Delete(id)
	if err != nil {
		log.Printf("UserService: %s", err)
		return err
	}

	return nil
}
