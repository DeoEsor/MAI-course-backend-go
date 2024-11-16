package pet

import (
	"fmt"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/passport"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
	"github.com/google/uuid"
	"time"
)

type Pet struct {
	ID        uuid.UUID `json:"id"`
	Status    status.Status
	Name      string
	Passport  *passport.Passport
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (pet *Pet) String() string {
	return fmt.Sprintf(
		"pet with id %v on state %v with name %v created at %v update at %v",
		pet.ID,
		pet.Status,
		pet.Name,
		pet.CreatedAt,
		pet.UpdatedAt,
	)
}
