package http

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"maker-checker/logger"
	"net/http"
)

// MakeHTTPHandler mounts all the API endpoints into a http.Handler.
func MakeHTTPHandler(log logger.ContextLogger, api ApiHandler) http.Handler {
	r := mux.NewRouter()
	e := MakeServerEndpoints(api)

	r.Methods("GET").Path("/health").Handler(http.HandlerFunc(healthHandler))

	r.Methods("POST").Path("/profiles/").Handler(httptransport.NewServer(
		e.PostProfileEndpoint,
		decodePostProfileRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/profiles/{id}").Handler(httptransport.NewServer(
		e.GetProfileEndpoint,
		decodeGetProfileRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/documents/{id}").Handler(httptransport.NewServer(
		e.GetDocumentEndpoint,
		decodeGetDocumentRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/documents/").Handler(httptransport.NewServer(
		e.PostDocumentEndpoint,
		decodePostDocumentRequest,
		encodeResponse,
	))

	r.Methods("POST").Path("/documents/requests/").Handler(httptransport.NewServer(
		e.PostDocumentRequestEndpoint,
		decodePostDocumentRequestRequest,
		encodeResponse,
	))

	r.Methods("PUT").Path("/documents/requests/{id}").Handler(httptransport.NewServer(
		e.PutDocumentRequestEndpoint,
		decodePutDocumentRequestRequest,
		encodeResponse,
	))

	r.Methods("GET").Path("/documents/requests/{id}").Handler(httptransport.NewServer(
		e.GetDocumentRequestEndpoint,
		decodeGetDocumentRequestRequest,
		encodeResponse,
	))

	return r
}

func healthHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	okMsg := []byte(`{"status": "ok"}`)
	_, err := writer.Write(okMsg)
	if err != nil {
		return
	}
}

func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	e, ok := response.(DomainError)
	if ok && e.error() != nil {
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(w).Encode(response)
}
