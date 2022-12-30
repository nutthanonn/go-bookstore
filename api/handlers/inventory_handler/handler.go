package inventoryhandler

import "github.com/google/uuid"

type inventoryHandler struct {
}

func (i *inventoryHandler) IsValidUUID(u string) bool {
	_, err := uuid.Parse(u)
	return err == nil
}

func NewInventoryHandler() *inventoryHandler {
	return &inventoryHandler{}
}
