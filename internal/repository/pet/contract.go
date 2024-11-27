package pet

import (
	"context"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet"
	"time"

	"github.com/google/uuid"
)

//go:generate mockgen -source=./contract.go -destination=./contract.mock.gen.go -package=model_run -typed=true

type RepositoryCreate interface {
	Create(ctx context.Context, pet *pet.Pet) error
}

type RepositoryGetter interface {
	Get(ctx context.Context, petId uuid.UUID) (*pet.Pet, error)
	GetRequired(ctx context.Context, petId uuid.UUID) (*pet.Pet, error)
}

type RepositorySearcher interface {
	Search(ctx context.Context, searchQuery *RepositorySearchQuery) ([]*pet.Pet, error)
}

type RepositoryUpdater interface {
	Update(ctx context.Context, modelRun *pet.Pet, version time.Time) error
}

type RepositoryI interface {
	RepositoryCreate
	RepositoryGetter
	RepositoryUpdater
	RepositorySearcher
}

type RepositorySearchQuery struct {
	Name      *string    `db:"name"`
	ProjectId *uuid.UUID `db:"project_id"`
	Offset    int        `db:"offset"`
	Limit     int        `db:"limit"`
}
