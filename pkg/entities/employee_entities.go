package entities

import (
	"time"

	"github.com/google/uuid"
)

type Employees struct {
	Employee_id   uuid.UUID `gorm:"primary_key;type:varchar(255)" json:"employee_id"`
	First_name    string    `gorm:"not null" json:"first_name"`
	Last_name     string    `gorm:"not null" json:"last_name"`
	Email         string    `gorm:"not null;unique" json:"email"`
	Phone         string    `gorm:"not null" json:"phone"`
	Salary        int       `gorm:"default:10000" json:"salary"`
	Job_title     string    `gorm:"not null" json:"job_title"`
	Date_of_birth string    `gorm:"not null" json:"date_of_birth"`
	Hire_date     time.Time `gorm:"autoCreateTime" json:"hire_date"`
	Create_at     time.Time `gorm:"autoCreateTime" json:"create_at"`
	Update_at     time.Time `gorm:"autoUpdateTime" json:"update_at"`

	// foreign key for Sales table
	Sales []Sales `gorm:"foreignkey:Employee_id"`
}

func (Employees) TableName() string { return "employees" }
