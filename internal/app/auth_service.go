package app

import (
	"errors"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/database"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/upper/db/v4"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService interface {
	Register(user domain.User) (domain.User, string, error)
	Login(user domain.User) (domain.User, string, error)
	Logout(sess domain.Session) error
	Check(sess domain.Session) error
	GenerateJwt(user domain.User) (string, error)
}

type authService struct {
	authRepo  database.SessionRepository
	userRepo  database.UserRepository
	tokenAuth *jwtauth.JWTAuth
	jwtTTL    time.Duration
}

func NewAuthService(ar database.SessionRepository, ur database.UserRepository, ta *jwtauth.JWTAuth, jwtTtl time.Duration) AuthService {
	return authService{
		authRepo:  ar,
		userRepo:  ur,
		tokenAuth: ta,
		jwtTTL:    jwtTtl,
	}
}

func (s authService) Register(user domain.User) (domain.User, string, error) {
	_, err := s.userRepo.FindByEmail(user.Email)
	if err == nil {
		log.Printf("invalid credentials")
		return domain.User{}, "", errors.New("invalid credentials")
	} else if !errors.Is(err, db.ErrNoMoreRows) {
		log.Print(err)
		return domain.User{}, "", err
	}

	user.Password, err = s.generatePasswordHash(user.Password)
	if err != nil {
		log.Printf("UserService: %s", err)
		return domain.User{}, "", err
	}

	user, err = s.userRepo.Save(user)
	if err != nil {
		log.Print(err)
		return domain.User{}, "", err
	}

	token, err := s.GenerateJwt(user)
	return user, token, err
}

func (s authService) Login(user domain.User) (domain.User, string, error) {
	u, err := s.userRepo.FindByEmail(user.Email)
	if err != nil {
		if errors.Is(err, db.ErrNoMoreRows) {
			log.Printf("AuthService: failed to find user %s", err)
		}
		log.Printf("AuthService: login error %s", err)
		return domain.User{}, "", err
	}

	valid := s.checkPasswordHash(user.Password, u.Password)
	if !valid {
		return domain.User{}, "", errors.New("invalid credentials")
	}

	token, err := s.GenerateJwt(u)
	if err != nil {
		log.Printf("AuthService->s.GenerateJwt %s", err)
		return domain.User{}, "", err
	}

	return u, token, err
}

func (s authService) Logout(sess domain.Session) error {
	return s.authRepo.Delete(sess)
}

func (s authService) GenerateJwt(user domain.User) (string, error) {
	sess := domain.Session{UserId: user.Id, UUID: uuid.New()}
	err := s.authRepo.Save(sess)
	if err != nil {
		log.Printf("AuthService: failed to save session %s", err)
		return "", err
	}

	claims := map[string]interface{}{
		"user_id": sess.UserId,
		"uuid":    sess.UUID,
	}
	jwtauth.SetExpiryIn(claims, s.jwtTTL)
	_, tokenString, err := s.tokenAuth.Encode(claims)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (s authService) Check(sess domain.Session) error {
	return s.authRepo.Exists(sess)
}

func (s authService) generatePasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (s authService) checkPasswordHash(password, hash string) bool {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) == nil
}
