package models

import (
	"fmt"

	"github.com/google/uuid"
)

type ErrModelNotFound struct {
	ModelName string
	Id        uuid.UUID
}

func (e *ErrModelNotFound) Error() string {
	return fmt.Sprintf("%s with Id [%d] does not exist.", e.ModelName, e.Id)
}
