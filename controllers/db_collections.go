package controllers

import "go.mongodb.org/mongo-driver/mongo"

var (
    EmployeeCol   *mongo.Collection
    AttendanceCol *mongo.Collection
    PayrollCol    *mongo.Collection
)

