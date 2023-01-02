package entities

import "time"

type User struct {
	ID         string    `gorm:"primary_key" json:"id"`
	FirstName  string    `gorm:"not null" json:"first_name"`
	LastName   string    `gorm:"not null" json:"last_name"`
	Created_at time.Time `json:"created_at"`
	Updated_at time.Time `json:"updated_at"`
}

func (User) TableName() string { return "users" }
