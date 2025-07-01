package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Payroll struct {
	ID         primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID primitive.ObjectID `bson:"employee_id" json:"employee_id"`
	Month      int                `bson:"month" json:"month"`
	Year       int                `bson:"year" json:"year"`
	TotalHours float64           `bson:"total_hours" json:"total_hours"`
	TotalPay   float64           `bson:"total_pay" json:"total_pay"`
	GeneratedAt time.Time        `bson:"generated_at" json:"generated_at"`
}
