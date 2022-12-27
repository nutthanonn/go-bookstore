package entities

import (
	"time"

	"github.com/google/uuid"
)

type Customers struct {
	Customer_id uuid.UUID `gorm:"primary_key;type:varchar(255)" json:"customer_id"`
	First_name  string    `json:"first_name"`
	Last_name   string    `json:"last_name"`
	Email       string    `json:"email"`
	Phone       string    `json:"phone"`
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
