package http

import (
	"context"
	"github.com/jmoiron/sqlx"
	"maker-checker/logger"
	doc "maker-checker/pkg/document"
	docRequest "maker-checker/pkg/document/request"
	prof "maker-checker/pkg/profile"
	"maker-checker/store/psql/document"
	"maker-checker/store/psql/document/request"
	"maker-checker/store/psql/profile"
)

// ApiHandler is the interface that wraps the basic API methods.
type ApiHandler interface {
	PostProfile(ctx context.Context, p Profile) (string, error)
	GetProfile(ctx context.Context, id string) (Profile, error)
	PostDocument(ctx context.Context, d Document) (string, error)
	GetDocument(ctx context.Context, id string) (Document, error)
	PostDocumentRequest(ctx context.Context, d DocumentRequest) (string, error)
	PutDocumentRequest(ctx context.Context, id string, d DocumentRequest) (string, error)
	GetDocumentRequest(ctx context.Context, id string) (DocumentRequest, error)
}

// Handler is the implementation of the ApiHandler interface.
type Handler struct {
	log           logger.ContextLogger
	profileSvc    *prof.Service
	documentSvc   *doc.Service
	docRequestSvc *docRequest.Service
}

// NewHandler returns a new instance of the Handler struct.
func NewHandler(log logger.ContextLogger, db *sqlx.DB) ApiHandler {
	profileSvc := prof.NewService(log, profile.NewRepository(log, db))
	documentSvc := doc.NewService(log, document.NewRepository(log, db))
	docRequestSvc := docRequest.NewService(log, request.NewRepository(log, db))
	return &Handler{
		log:           log,
		profileSvc:    profileSvc,
		documentSvc:   documentSvc,
		docRequestSvc: docRequestSvc,
	}
}

func (h *Handler) PostProfile(ctx context.Context, payload Profile) (string, error) {
	id, err := h.profileSvc.CreateProfile(ctx, prof.Row{
		FirstName: payload.FirstName,
		LastName:  payload.LastName,
		Email:     payload.Email,
	})

	if err != nil {
		return "", ErrFailedToCreate
	}

	return id, nil
}

func (h *Handler) GetProfile(ctx context.Context, id string) (Profile, error) {
	p, err := h.profileSvc.FindByID(ctx, id)
	if err != nil {
		return Profile{}, ErrNotFound
	}
	res := Profile{
		ID:        p.ID,
		FirstName: p.FirstName,
		LastName:  p.LastName,
		Email:     p.Email,
	}
	return res, nil
}

func (h *Handler) PostDocument(ctx context.Context, payload Document) (string, error) {
	id, err := h.documentSvc.CreateDocument(ctx, doc.Row{
		Description: payload.Description,
		CreatorID:   payload.CreatorID,
		Status:      payload.Status,
	})

	if err != nil {
		return "", ErrFailedToCreate
	}

	return id, nil
}

func (h *Handler) GetDocument(ctx context.Context, id string) (Document, error) {
	d, err := h.documentSvc.FindByID(ctx, id)
	if err != nil {
		return Document{}, ErrNotFound
	}
	result := Document{
		ID:          d.ID,
		Description: d.Description,
		CreatorID:   d.CreatorID,
		Status:      d.Status,
	}
	return result, nil
}

func (h *Handler) PostDocumentRequest(ctx context.Context, payload DocumentRequest) (string, error) {
	approvers := make([]*docRequest.Approver, len(payload.Approvers))
	for i, approver := range payload.Approvers {
		approvers[i] = &docRequest.Approver{
			ID: approver.ID,
		}
	}

	id, err := h.docRequestSvc.CreateDocumentRequest(ctx, docRequest.Row{
		DocumentID:     payload.DocumentID,
		CreatorID:      payload.CreatorID,
		Approvers:      approvers,
		ApproverCount:  payload.ApproverCount,
		RecipientEmail: payload.RecipientEmail,
	})

	if err != nil {
		return "", ErrFailedToCreate
	}

	return id, nil
}

func (h *Handler) PutDocumentRequest(ctx context.Context, reqID string, payload DocumentRequest) (string, error) {
	if reqID == "" {
		return "", ErrInconsistentIDs
	}

	approvers := make([]*docRequest.Approver, len(payload.Approvers))
	for i, approver := range payload.Approvers {
		approvers[i] = &docRequest.Approver{
			ID:     approver.ID,
			Status: approver.Status,
		}
	}

	id, err := h.docRequestSvc.UpdateDocumentRequest(ctx, docRequest.Row{
		ID:            reqID,
		DocumentID:    payload.DocumentID,
		CreatorID:     payload.CreatorID,
		Approvers:     approvers,
		ApproverCount: payload.ApproverCount,
	})
	if err != nil {
		return "", ErrFailedToUpdate
	}

	return id, nil
}

func (h *Handler) GetDocumentRequest(ctx context.Context, id string) (DocumentRequest, error) {
	d, err := h.docRequestSvc.FindByID(ctx, id)
	if err != nil {
		return DocumentRequest{}, ErrNotFound
	}

	approvers := make([]RequestApprover, len(d.Approvers))
	for i, approver := range d.Approvers {
		approvers[i] = RequestApprover{
			ID:     approver.ID,
			Status: approver.Status,
		}
	}

	res := DocumentRequest{
		ID:             d.ID,
		DocumentID:     d.DocumentID,
		CreatorID:      d.CreatorID,
		Approvers:      approvers,
		ApproverCount:  d.ApproverCount,
		Status:         d.Status,
		RecipientEmail: d.RecipientEmail,
	}
	return res, nil
}
