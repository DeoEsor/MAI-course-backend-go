package update

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/passport"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
	"github.com/google/uuid"
)

type Command struct {
	PetId    uuid.UUID
	Name     string
	Passport *passport.Passport
	Status   status.Status
}

type CommandResponse struct {
}

type CommandHandler interface {
	Update(ctx context.Context, command *Command) (*CommandResponse, error)
}
