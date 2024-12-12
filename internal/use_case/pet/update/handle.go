package update

import (
	"context"
	log "github.com/sirupsen/logrus"
	"time"
)

func (useCase *useCase) Update(ctx context.Context, command *Command) (*CommandResponse, error) {
	pet, err := useCase.petRepository.GetRequired(ctx, command.PetId)
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			Error("failed to get pet")
		return nil, err
	}
	version := pet.UpdatedAt

	pet.Name = command.Name
	pet.Passport = command.Passport
	pet.Status = command.Status
	pet.UpdatedAt = time.Now().UTC()

	if err := useCase.petRepository.Update(ctx, pet, version); err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithField("pet_id", pet.ID).
			Error("failed to update pet")
		return nil, err
	}

	return &CommandResponse{}, nil
}
