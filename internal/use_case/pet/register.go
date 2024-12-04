package pet

import (
	"github.com/DeoEsor/MAI-course-backend-go/internal/use_case/pet/create"
	"github.com/samber/do/v2"
)

func Register(injector do.Injector) error {

	do.Provide(injector, create.New)

	return nil
}
