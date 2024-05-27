package middlewares

import (
	"context"
	"errors"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/app"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/domain"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
	"github.com/lestrrat-go/jwx/v2/jwt"
	"github.com/upper/db/v4"
	"net/http"
)

func AuthMiddleware(ja *jwtauth.JWTAuth, as app.AuthService, us app.UserService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token, err := jwtauth.VerifyRequest(ja, r, jwtauth.TokenFromHeader)

			if err != nil {
				controllers.Unauthorized(w, err)
				return
			}

			if token == nil || jwt.Validate(token) != nil {
				controllers.Unauthorized(w, err)
				return
			}

			claims := token.PrivateClaims()
			uId := uint64(claims["user_id"].(float64))
			uUuid, err := uuid.Parse(claims["uuid"].(string))
			if err != nil {
				controllers.Unauthorized(w, err)
				return
			}

			auth := domain.Session{
				UserId: uId,
				UUID:   uUuid,
			}
			err = as.Check(auth)
			if err != nil {
				controllers.Unauthorized(w, err)
				return
			}

			user, err := us.FindById(uId)
			if err != nil {
				if errors.Is(err, db.ErrNoMoreRows) {
					err = errors.New("unauthorized")
				}
				controllers.Unauthorized(w, err)
				return
			}

			ctx = context.WithValue(ctx, controllers.UserKey, user)
			ctx = context.WithValue(ctx, controllers.SessKey, auth)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}
