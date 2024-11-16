package pet

import (
	"fmt"
	"github.com/DeoEsor/MAI-course-backend-go/internal/domain/pet/value_object/status"
	"time"
)

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
