package entities

import (
	"time"

	"github.com/google/uuid"
)

type Customers struct {
	Customer_id uuid.UUID `gorm:"primary_key;type:varchar(255)" json:"customer_id"`
	First_name  string    `gorm:"not null" json:"first_name"`
	Last_name   string    `gorm:"not null" json:"last_name"`
	Email       string    `gorm:"not null;unique" json:"email"`
	Phone       string    `gorm:"not null" json:"phone"`
	Address     string    `json:"address"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Zip         string    `json:"zip"`
	Create_at   time.Time `json:"create_at"`
	Update_at   time.Time `json:"update_at"`

	// foreign key for Orders table
	Orders []Orders `gorm:"foreignkey:Customer_id"`
}

func (Customers) TableName() string { return "customers" }
