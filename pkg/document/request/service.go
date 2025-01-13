package request

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"maker-checker/logger"
	helpers "maker-checker/pkg"
	"time"
)

type Service struct {
	log  logger.ContextLogger
	repo Repository
}

func NewService(log logger.ContextLogger, repo Repository) *Service {
	return &Service{log: log, repo: repo}
}

func (s *Service) FindByID(ctx context.Context, id string) (Row, error) {
	result, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return Row{}, errors.New("failed to find document request by id")
	}

	return *result, nil
}

func (s *Service) CreateDocumentRequest(ctx context.Context, input Row) (string, error) {
	requestID := uuid.NewString()
	input.ID = requestID

	for _, approver := range input.Approvers {
		approver.Status = Pending
	}
	now := time.Now().UTC()
	input.Status = Pending

	// Set the default number of approvers to 2 if not provided
	if input.ApproverCount == 0 {
		input.ApproverCount = 2
	}

	input.CreatedAt = now
	input.UpdatedAt = now

	id, err := s.repo.Create(ctx, input)
	if err != nil {
		return "", errors.New("failed to create document request")
	}

	return id, nil
}

func (s *Service) UpdateDocumentRequest(ctx context.Context, input Row) (string, error) {
	docRequest, err := s.repo.FindByID(ctx, input.ID)
	if err != nil {
		s.log.Errorw(ctx, "failed to find document request by id", "id", input.ID)
		return "", err
	}

	now := time.Now().UTC()

	for _, approver := range input.Approvers {
		if approver.Status == Approved {
			approver.ApprovedAt = now.Format("2006-01-02T15:04:05Z")
		}
	}

	if len(input.Approvers) == docRequest.ApproverCount && helpers.Every(input.Approvers, func(approver *Approver) bool {
		return approver.Status == Approved
	}) {
		input.Status = Approved
	}

	if helpers.Some(input.Approvers, func(approver *Approver) bool {
		return approver.Status == Rejected
	}) {
		input.Status = Rejected
	}

	input.UpdatedAt = now

	id, err := s.repo.Update(ctx, input)
	if err != nil {
		return "", errors.New("failed to update document request")
	}
	// Notify the recipient or creator of the document request
	_ = s.Notify(ctx, input)

	return id, nil
}

func (s *Service) Notify(ctx context.Context, input Row) error {
	if input.Status == Approved {
		s.log.Infow(ctx, "notifying recipient", "id", input.ID)
		// TODO: Publish message to the message broker to notify the recipient that the document request has been approved

	}
	if input.Status == Rejected {
		s.log.Infow(ctx, "notifying creator", "id", input.ID)
		//	TODO: Publish message to the message broker to notify the creator that the document request has been rejected
	}
	return nil
}
