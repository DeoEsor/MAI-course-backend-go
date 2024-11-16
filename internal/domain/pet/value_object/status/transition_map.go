package status

import (
	"fmt"
	"slices"
)

var (
	transitionMap = map[Status][]Status{
		OnMedicalExamination: []Status{
			OnRecover,
		},
		OnRecover: {
			WaitingOwner,
		},
		WaitingOwner: {
			Adopted,
		},
		Adopted: {}, // конечное значение
	}
)

func (state Status) IsTransitionAllowed(newStatus Status) (bool, error) {
	allowedToStatuses, ok := transitionMap[state]
	if !ok {
		return false, fmt.Errorf("unknown transition state: %s", state)
	}

	isAllowedTo := slices.Contains(allowedToStatuses, newStatus)

	return isAllowedTo, nil
}
