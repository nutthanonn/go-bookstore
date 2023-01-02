package entities

import (
	"time"

	"github.com/google/uuid"
)

type Sales struct {
	Order_id    uuid.UUID `gorm:"primary_key;" json:"order_id"`
	Employee_id uuid.UUID `json:"employee_id"`
	Create_at   time.Time `gorm:"autoCreateTime" json:"create_at"`
}
