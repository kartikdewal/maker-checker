package profile

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
		return Row{}, errors.New("failed to find profile by id")
	}

	return *result, nil
}

func (s *Service) CreateProfile(ctx context.Context, input Row) (string, error) {
	profileID := uuid.NewString()
	input.ID = profileID

	now := time.Now().UTC()

	input.CreatedAt = now
	input.UpdatedAt = now

	id, err := s.repo.Create(ctx, input)
	if err != nil {
		return "", errors.New("failed to create profile")
	}

	return id, nil
}
