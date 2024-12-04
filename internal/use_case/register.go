package use_case

import (
	"errors"
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case/pet"
	"github.com/samber/do/v2"
)

func Register(injector do.Injector) error {

	return errors.Join(
		pet.Register(injector),
	)
}
