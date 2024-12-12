package get

import (
	"context"
	log "github.com/sirupsen/logrus"
)

func (useCase *useCase) Get(ctx context.Context, command *Command) (*CommandResponse, error) {

	pet, err := useCase.petRepository.GetRequired(ctx, command.PetId)
	if err != nil {
		log.
			WithContext(ctx).
			WithError(err).
			WithField("pet_id", command.PetId).
			Error("failed to get pet")
		return nil, err
	}

	return &CommandResponse{
		Pet: pet,
	}, nil
}
