package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Attendance struct {
	ID         primitive.ObjectID  `bson:"_id,omitempty" json:"id,omitempty"`
	EmployeeID primitive.ObjectID  `bson:"employee_id" json:"employee_id"`
	CheckIn    time.Time           `bson:"check_in" json:"check_in"`
	CheckOut   *time.Time          `bson:"check_out,omitempty" json:"check_out,omitempty"`
	WorkHours  float64             `bson:"work_hours" json:"work_hours"`
	Date       time.Time           `bson:"date" json:"date"`
}
