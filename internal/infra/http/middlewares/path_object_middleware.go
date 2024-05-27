package middlewares

import (
	"context"
	"fmt"
	"github.com/BohdanBoriak/boilerplate-go-back/internal/infra/http/controllers"
	"github.com/go-chi/chi/v5"
	"github.com/upper/db/v4"
	"log"
	"net/http"
	"strconv"
)

type Findable interface {
	Find(uint64) (interface{}, error)
}

func PathObject(pathKey string, ctxKey controllers.CtxKey, service Findable) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		hfn := func(w http.ResponseWriter, r *http.Request) {
			id, err := strconv.ParseUint(chi.URLParam(r, pathKey), 10, 64)
			if err != nil {
				err = fmt.Errorf("invalid %s parameter(only non-negative integers)", pathKey)
				log.Print(err)
				controllers.BadRequest(w, err)
				return
			}

			obj, err := service.Find(id)
			if err != nil {
				log.Print(err)
				errInt4 := fmt.Errorf("%d is greater than maximum value for Int4", id)
				if err == db.ErrNoMoreRows || err.Error() == errInt4.Error() {
					err = fmt.Errorf("record not found")
					controllers.NotFound(w, err)
					return
				}
				controllers.InternalServerError(w, err)
				return
			}

			ctx := context.WithValue(r.Context(), ctxKey, obj)

			next.ServeHTTP(w, r.WithContext(ctx))
		}
		return http.HandlerFunc(hfn)
	}
}
