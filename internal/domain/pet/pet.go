package pet

import (
	"fmt"
	"time"

	"github.com/google/uuid"

	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/passport"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
)

type Pet struct {
	ID        uuid.UUID          `json:"id"`
	Status    status.Status      `json:"status"`
	Name      string             `json:"name"`
	Passport  *passport.Passport `json:"passport"`
	CreatedAt time.Time          `json:"created_at"`
	UpdatedAt time.Time          `json:"updated_at"`
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
