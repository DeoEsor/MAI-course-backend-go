package create

import (
	"context"
	pet2 "github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	log "github.com/sirupsen/logrus"
)

func (useCase *useCase) Create(ctx context.Context, command *Command) (*CommandResponse, error) {
	pet := pet2.New(command.Name)

	err := useCase.petRepository.Create(ctx, pet)
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			Error("failed to create pet")
		return nil, err
	}

	return &CommandResponse{
		PetId: pet.ID,
	}, nil
}
