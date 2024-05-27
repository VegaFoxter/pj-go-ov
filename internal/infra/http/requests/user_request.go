package requests

import (
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
)

type RegisterRequest struct {
	FirstName  string `json:"firstName" validate:"required,gte=1,max=40"`
	SecondName string `json:"secondName" validate:"required,gte=1,max=40"`
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,gte=4,max=20"`
}

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password"  validate:"required,gte=4"`
}

type UpdateUserRequest struct {
	FirstName  string `json:"firstName" validate:"required,gte=1,max=40"`
	SecondName string `json:"secondName" validate:"required,gte=1,max=40"`
	Email      string `json:"email" validate:"required,email"`
}

func (r RegisterRequest) ToDomainModel() (interface{}, error) {
	return domain.User{
		FirstName:  r.FirstName,
		SecondName: r.SecondName,
		Email:      r.Email,
		Password:   r.Password,
	}, nil
}

func (r UpdateUserRequest) ToDomainModel() (interface{}, error) {
	return domain.User{
		FirstName:  r.FirstName,
		SecondName: r.SecondName,
		Email:      r.Email,
	}, nil
}

func (r LoginRequest) ToDomainModel() (interface{}, error) {
	return domain.User{
		Password: r.Password,
		Email:    r.Email,
	}, nil
}
