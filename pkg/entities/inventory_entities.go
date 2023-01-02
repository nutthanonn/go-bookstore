package entities

import (
	"time"

	"github.com/google/uuid"
)

type Inventories struct {
	Book_id    uuid.UUID `gorm:"primary_key" json:"book_id"`
	Quantity   int       `gorm:"default:1" json:"quantity"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (Inventories) TableName() string { return "inventories" }
