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
	Passport  passport.Passport
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (pet *Pet) UpdateName(newName string) {
	pet.Name = newName
	pet.UpdatedAt = time.Now().UTC()
}

func (pet *Pet) UpdateStatus(newStatus status.Status) error {
	allowed, err := pet.Status.IsTransitionAllowed(newStatus)
	if err != nil {
		return err
	}
	if !allowed {
		return fmt.Errorf("pet %v is not allowed to change status [from state %v, to state %v]", pet.ID, pet.Status, newStatus)
	}

	pet.Status = newStatus
	pet.UpdatedAt = time.Now().UTC()
	return nil
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
