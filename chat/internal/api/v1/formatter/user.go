package formatter

import (
	"fmt"
	"strings"

	"github.com/google/uuid"
)

type DomainIDType string

func (uid *DomainIDType) ToDomain() (uuid.UUID, error) {
	userStringID := strings.TrimSpace(string(*uid))
	if userStringID == "" {
		return uuid.UUID{}, fmt.Errorf("%w %s", ErrInvalidData, "Поле <ИД> не может быть пустым")
	}

	userID, err := uuid.Parse(userStringID)

	if err != nil {
		return uuid.UUID{}, fmt.Errorf("%w %s", ErrInvalidData, "Поле <ИД> не корректно")
	}

	return userID, nil
}
