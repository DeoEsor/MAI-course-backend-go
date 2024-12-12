package get

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"github.com/google/uuid"
)

type Command struct {
	PetId uuid.UUID
}

type CommandResponse struct {
	Pet *pet.Pet
}

type CommandHandler interface {
	Get(ctx context.Context, command *Command) (*CommandResponse, error)
}
