package create

type UseCase struct {
	// зависимости use-case
	// репозиторий для работы с БД
	// стороние сервисы и т.д.
	Name string
}

func New() (Handler, error) {
	return &UseCase{}, nil
}
