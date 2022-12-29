package entities

import (
	"time"

	"github.com/google/uuid"
)

type Books struct {
	Book_id      uuid.UUID `gorm:"primary_key" json:"book_id"`
	Title        string    `json:"title"`
	Author       string    `json:"author"`
	Publish_year string    `json:"publish_year"`
	Price        float32   `json:"price"`
	Genre        string    `json:"genre"`
	Created_at   time.Time `json:"created_at"`
	Updated_at   time.Time `json:"updated_at"`

	// foreign key for Inventory table
	Inventory_id Inventories `gorm:"foreignkey:Book_id"`

	// foreign key for OrderDetails table
	OrderDetails []OrderDetails `gorm:"foreignkey:Book_id"`
}

func (Books) TableName() string { return "books" }
