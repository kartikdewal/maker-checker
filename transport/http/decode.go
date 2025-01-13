package http

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"net/http"
)

var (
	ErrBadRouting = errors.New("inconsistent mapping between route and handler")
)

func decodeGetProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetProfileRequest{ID: id}, nil
}

func decodePostProfileRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req PostProfileRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Profile); e != nil {
		return nil, e
	}
	return req, nil
}

func decodeGetDocumentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetDocumentRequest{ID: id}, nil
}

func decodePostDocumentRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req PostDocumentRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Document); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePostDocumentRequestRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	var req PostDocumentRequestRequest
	if e := json.NewDecoder(r.Body).Decode(&req.DocumentRequest); e != nil {
		return nil, e
	}
	return req, nil
}

func decodePutDocumentRequestRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	var req DocumentRequest
	if e := json.NewDecoder(r.Body).Decode(&req); e != nil {
		return nil, e
	}
	return PutDocumentRequestRequest{
		ID:              id,
		DocumentRequest: req,
	}, nil
}

func decodeGetDocumentRequestRequest(_ context.Context, r *http.Request) (request interface{}, err error) {
	vars := mux.Vars(r)
	id, ok := vars["id"]
	if !ok {
		return nil, ErrBadRouting
	}
	return GetDocumentRequestRequest{ID: id}, nil
}
