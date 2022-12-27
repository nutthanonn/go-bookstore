package entities

import (
	"time"

	"github.com/google/uuid"
)

type Employees struct {
	Employee_id   uuid.UUID `gorm:"primary_key;type:varchar(255)" json:"employee_id"`
	First_name    string    `json:"first_name"`
	Last_name     string    `json:"last_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Salary        int       `json:"salary"`
	Job_title     string    `json:"job_title"`
	Date_of_birth time.Time `json:"date_of_birth"`
	Hire_date     time.Time `json:"hire_date"`
	Create_at     time.Time `json:"create_at"`
	Update_at     time.Time `json:"update_at"`

	// foreign key for Sales table
	Sales []Sales `gorm:"foreignkey:Employee_id"`
}

func (Employees) TableName() string { return "employees" }
