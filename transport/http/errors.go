package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
)

var (
	ErrInconsistentIDs = errors.New("inconsistent IDs")
	ErrAlreadyExists   = errors.New("already exists")
	ErrNotFound        = errors.New("not found")
	ErrFailedToCreate  = errors.New("failed to create")
	ErrFailedToUpdate  = errors.New("failed to update")
)

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(getErrorCode(err))
	_ = json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func getErrorCode(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrAlreadyExists), errors.Is(err, ErrInconsistentIDs):
		return http.StatusBadRequest
	default:
		return http.StatusInternalServerError
	}
}

type DomainError interface {
	error() error
}
