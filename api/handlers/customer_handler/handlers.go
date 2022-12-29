package customerhandler

import "github.com/google/uuid"

type customerHandlers struct {
}

func (e *customerHandlers) IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func NewCustomerHandlers() *customerHandlers {
	return &customerHandlers{}
}
