package search

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/repository/pet"
	log "github.com/sirupsen/logrus"
)

func (useCase *useCase) Search(ctx context.Context, command *Command) (*CommandResponse, error) {
	if command.Name != nil {
		nameSearch := "%" + *command.Name + "%"
		command.Name = &nameSearch
	}

	query := &pet.RepositorySearchQuery{
		Name:   command.Name,
		Offset: command.Offset,
		Limit:  command.Limit,
	}

	pets, err := useCase.petRepository.Search(ctx, query)
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithField("query", query).
			Error("failed to find pets")
		return nil, err
	}

	return &CommandResponse{
		Pets: pets,
	}, nil
}
