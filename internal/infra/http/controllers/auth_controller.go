package controllers

import (
	"errors"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/requests"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/resources"
	"log"
	"net/http"
)

type AuthController struct {
	authService app.AuthService
	userService app.UserService
}

func NewAuthController(as app.AuthService, us app.UserService) AuthController {
	return AuthController{
		authService: as,
		userService: us,
	}
}

func (c AuthController) Register() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := requests.Bind(r, requests.RegisterRequest{}, domain.User{})
		if err != nil {
			log.Printf("AuthController: %s", err)
			BadRequest(w, errors.New("invalid request body"))
			return
		}

		user, token, err := c.authService.Register(user)
		if err != nil {
			log.Printf("AuthController: %s", err)
			BadRequest(w, err)
			return
		}

		var authDto resources.AuthDto
		Success(w, authDto.DomainToDto(token, user))
	}
}

func (c AuthController) Login() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := requests.Bind(r, requests.LoginRequest{}, domain.User{})
		if err != nil {
			log.Printf("AuthController: %s", err)
			BadRequest(w, err)
			return
		}

		u, token, err := c.authService.Login(user)
		if err != nil {
			log.Printf("AuthController: %s", err)
			InternalServerError(w, err)
			return
		}

		var authDto resources.AuthDto
		Success(w, authDto.DomainToDto(token, u))
	}
}

func (c AuthController) Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess := r.Context().Value(SessKey).(domain.Session)
		err := c.authService.Logout(sess)
		if err != nil {
			log.Printf("AuthController: %s", err)
			InternalServerError(w, err)
			return
		}

		noContent(w)
	}
}
