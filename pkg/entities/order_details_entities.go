package entities

import "github.com/google/uuid"

type OrderDetails struct {
	Order_id uuid.UUID `gorm:"primary_key;" json:"order_id"`
	Book_id  uuid.UUID `gorm:"primary_key;" json:"book_id"`
	Quantity int       `gorm:"not null;default:0" json:"quantity"`
	Coast    float32   `json:"coast"`
}

func (OrderDetails) TableName() string { return "order_details" }
