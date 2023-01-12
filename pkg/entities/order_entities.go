package entities

import (
	"time"

	"github.com/google/uuid"
)

type Orders struct {
	Order_id    uuid.UUID `gorm:"primary_key;" json:"order_id"`
	Customer_id uuid.UUID `json:"customer_id"`
	Create_at   time.Time `gorm:"autoCreateTime" json:"create_at"`
	Update_at   time.Time `gorm:"autoUpdateTime" json:"update_at"`
	Status      bool      `gorm:"default:false" json:"status"`

	// foreign key for OrderDetails table
	OrderDetails []OrderDetails `gorm:"foreignkey:Order_id" json:"order_details"`

	// foreign key for Sales table
	Sales Sales `gorm:"foreignkey:Order_id"`
}

func (Orders) TableName() string { return "orders" }
