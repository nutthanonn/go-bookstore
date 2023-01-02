package entities

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	Order_id    uuid.UUID `gorm:"primary_key;" json:"order_id"`
	Customer_id uuid.UUID `json:"customer_id"`
	Order_date  string    `gorm:"not null;default:null" json:"order_date"`
	Create_at   time.Time `json:"create_at"`
	Update_at   time.Time `json:"update_at"`

	// foreign key for OrderDetails table
	OrderDetails []OrderDetails `gorm:"foreignkey:Order_id"`

	// foreign key for Sales table
	Sales Sales `gorm:"foreignkey:Order_id"`
}

func (Orders) TableName() string { return "orders" }
