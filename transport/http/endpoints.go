package http

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"maker-checker/pkg/document/request"
)

// Endpoints collects all the endpoints that compose the API.
type Endpoints struct {
	PostProfileEndpoint         endpoint.Endpoint
	GetProfileEndpoint          endpoint.Endpoint
	GetDocumentEndpoint         endpoint.Endpoint
	PostDocumentEndpoint        endpoint.Endpoint
	PostDocumentRequestEndpoint endpoint.Endpoint
	PutDocumentRequestEndpoint  endpoint.Endpoint
	GetDocumentRequestEndpoint  endpoint.Endpoint
}

// MakeServerEndpoints returns an Endpoints struct where each endpoint invokes
// the corresponding method on the provided ApiHandler.
func MakeServerEndpoints(api ApiHandler) Endpoints {
	return Endpoints{
		PostProfileEndpoint:         MakePostProfileEndpoint(api),
		GetProfileEndpoint:          MakeGetProfileEndpoint(api),
		GetDocumentEndpoint:         MakeGetDocumentEndpoint(api),
		PostDocumentEndpoint:        MakePostDocumentEndpoint(api),
		PostDocumentRequestEndpoint: MakePostDocumentRequestEndpoint(api),
		PutDocumentRequestEndpoint:  MakePutDocumentRequestEndpoint(api),
		GetDocumentRequestEndpoint:  MakeGetDocumentRequestEndpoint(api),
	}
}

func MakeGetProfileEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetProfileRequest)
		p, e := api.GetProfile(ctx, req.ID)
		return GetProfileResponse{Profile: p, Err: e}, nil
	}
}

func MakePostProfileEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostProfileRequest)
		profileID, err := api.PostProfile(ctx, req.Profile)
		return PostProfileResponse{ProfileID: profileID, Err: err}, nil
	}
}

func MakeGetDocumentEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetDocumentRequest)
		d, e := api.GetDocument(ctx, req.ID)
		response = GetDocumentResponse{Document: d, Err: e}
		return response, nil
	}
}

func MakePostDocumentEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostDocumentRequest)
		id, e := api.PostDocument(ctx, req.Document)
		return PostDocumentResponse{Id: id, Err: e}, nil
	}
}

func MakePostDocumentRequestEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PostDocumentRequestRequest)
		id, e := api.PostDocumentRequest(ctx, req.DocumentRequest)
		return PostDocumentRequestResponse{Id: id, Err: e}, nil
	}
}

func MakePutDocumentRequestEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(PutDocumentRequestRequest)
		id, e := api.PutDocumentRequest(ctx, req.ID, req.DocumentRequest)
		return PutDocumentRequestResponse{Id: id, Err: e}, nil
	}
}

func MakeGetDocumentRequestEndpoint(api ApiHandler) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(GetDocumentRequestRequest)
		d, e := api.GetDocumentRequest(ctx, req.ID)
		return GetDocumentRequestResponse{DocumentRequest: d, Err: e}, nil
	}
}

type Profile struct {
	ID        string `json:"id,omitempty"`
	FirstName string `json:"firstName,omitempty"`
	LastName  string `json:"lastName,omitempty"`
	Email     string `json:"email,omitempty"`
}

type Document struct {
	ID          string `json:"id,omitempty"`
	Description string `json:"description,omitempty"`
	CreatorID   string `json:"creatorID,omitempty"`
	Status      string `json:"status,omitempty"`
}

type DocumentRequest struct {
	ID             string                 `json:"id,omitempty"`
	DocumentID     string                 `json:"documentID"`
	CreatorID      string                 `json:"creatorID"`
	Approvers      []RequestApprover      `json:"approvers"`
	ApproverCount  int                    `json:"approverCount"`
	Status         request.ApprovalStatus `json:"status,omitempty"`
	RecipientEmail string                 `json:"recipientEmail"`
}

type RequestApprover struct {
	ID     string                 `json:"id"`
	Status request.ApprovalStatus `json:"status"`
}
type GetProfileRequest struct {
	ID string
}

type PostProfileRequest struct {
	Profile Profile
}

type PostDocumentRequest struct {
	Document Document
}

type GetDocumentRequest struct {
	ID string
}

type PostDocumentRequestRequest struct {
	DocumentRequest DocumentRequest
}

type PutDocumentRequestRequest struct {
	ID              string
	DocumentRequest DocumentRequest
}

type GetDocumentRequestRequest struct {
	ID string
}

type GetProfileResponse struct {
	Profile Profile `json:"profile,omitempty"`
	Err     error   `json:"err,omitempty"`
}

func (r GetProfileResponse) error() error { return r.Err }

type PostProfileResponse struct {
	ProfileID string `json:"profile_id,omitempty"`
	Err       error  `json:"err,omitempty"`
}

func (r PostProfileResponse) error() error { return r.Err }

type GetDocumentResponse struct {
	Document Document `json:"document,omitempty"`
	Err      error    `json:"err,omitempty"`
}

func (r *GetDocumentResponse) error() error { return r.Err }

type PostDocumentResponse struct {
	Id  string `json:"id,omitempty"`
	Err error  `json:"err,omitempty"`
}

func (r PostDocumentResponse) error() error { return r.Err }

type PostDocumentRequestResponse struct {
	Id  string `json:"id,omitempty"`
	Err error  `json:"err,omitempty"`
}

func (r PostDocumentRequestResponse) error() error { return r.Err }

type PutDocumentRequestResponse struct {
	Id  string `json:"id,omitempty"`
	Err error  `json:"err,omitempty"`
}

type GetDocumentRequestResponse struct {
	DocumentRequest DocumentRequest `json:"document_request,omitempty"`
	Err             error           `json:"err,omitempty"`
}
