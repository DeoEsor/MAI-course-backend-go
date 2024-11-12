package passport

import "github.com/google/uuid"

type Passport struct {
	Id uuid.UUID `json:"id"`
}
