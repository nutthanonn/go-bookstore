package employeehandler

import "github.com/google/uuid"

type employeeHandlers struct {
}

func (e *employeeHandlers) IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func NewEmployeeHandlers() *employeeHandlers {
	return &employeeHandlers{}
}
