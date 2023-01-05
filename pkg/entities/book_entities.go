package entities

import (
	"time"

	"github.com/google/uuid"
)

type Books struct {
	Book_id      uuid.UUID `gorm:"primary_key" json:"book_id"`
	Title        string    `gorm:"not null" json:"title"`
	Author       string    `gorm:"not null" json:"author"`
	Publish_year string    `gorm:"not null" json:"publish_year"`
	Price        float32   `gorm:"not null" json:"price"`
	Genre        string    `json:"genre"`
	Created_at   time.Time `gorm:"autoCreateTime" json:"created_at"`
	Updated_at   time.Time `gorm:"autoUpdateTime" json:"updated_at"`

	// foreign key for Inventory table
	Inventory Inventories `gorm:"foreignkey:Book_id"`

	// foreign key for OrderDetails table
	OrderDetails []OrderDetails `gorm:"foreignkey:Book_id"`
}

func (Books) TableName() string { return "books" }
