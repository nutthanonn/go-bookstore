package bookhandler

import "github.com/google/uuid"

type bookHandler struct {
}

func (e *bookHandler) IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func NewBookHandler() *bookHandler {
	return &bookHandler{}
}
