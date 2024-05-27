package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

/* should not use built-in type string as key for value;
define your own type to avoid collisions */

type CtxKey struct {
	Name string
}

var (
	UserKey = CtxKey{Name: "user"}
	SessKey = CtxKey{Name: "sess"}
)

func Ok(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func Success(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Print(err)
	}
}

// nolint
func Created(w http.ResponseWriter, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	err := json.NewEncoder(w).Encode(body)
	if err != nil {
		log.Print(err)
	}
}

func noContent(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

func BadRequest(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	encodeErrorBody(w, err)
}

func Forbidden(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)

	encodeErrorBody(w, err)
}

func InternalServerError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusInternalServerError)

	encodeErrorBody(w, err)
}

// nolint
func validationError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnprocessableEntity)

	encodeErrorBody(w, err)
}

// nolint
func genericError(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)

	encodeErrorBody(w, err)
}

func NotFound(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)

	body := "Not Found"

	if err != nil {
		body = fmt.Sprint(err)
	}

	e := json.NewEncoder(w).Encode(map[string]interface{}{"error": body})
	if e != nil {
		log.Print(e)
	}
}

func encodeErrorBody(w http.ResponseWriter, err error) {
	e := json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
	if e != nil {
		log.Print(e)
	}
}

func Unauthorized(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusUnauthorized)

	encodeErrorBody(w, err)
}
