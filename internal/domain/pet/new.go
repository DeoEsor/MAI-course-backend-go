package pet

import (
	"time"

	"github.com/google/uuid"

	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/passport"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
)

func New(name string) *Pet {
	now := time.Now().UTC()
	id, _ := uuid.NewRandom()

	return &Pet{
		ID:        id,
		Status:    status.OnMedicalExamination,
		Name:      name,
		Passport:  &passport.Passport{},
		CreatedAt: now,
		UpdatedAt: now,
	}
}
