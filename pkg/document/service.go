package document

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"maker-checker/logger"
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
		s.log.Errorw(ctx, "failed to find document by id", "id", id)
		return Row{}, errors.New("failed to find document by id")
	}

	return *result, nil
}

func (s *Service) CreateDocument(ctx context.Context, input Row) (string, error) {
	documentID := uuid.NewString()
	input.ID = documentID

	now := time.Now().UTC()

	input.CreatedAt = now
	input.UpdatedAt = now

	id, err := s.repo.Create(ctx, input)
	if err != nil {
		return "", errors.New("failed to create document")
	}

	return id, nil
}
