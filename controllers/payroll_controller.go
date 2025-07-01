package controllers

import (
	"context"
	"net/http"
	"time"

	"hris-api/config"
	"hris-api/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

)

var payrollCol *mongo.Collection
// var employeeCol *mongo.Collection
// var attendanceCol *mongo.Collection

func InitPayrollController() {
	payrollCol = config.GetCollection("payrolls")
	// employeeCol = config.GetCollection("employees")
	// attendanceCol = config.GetCollection("attendances")
}

func GeneratePayroll(c *gin.Context) {
	var input struct {
		EmployeeID string `json:"employee_id"`
		Month      int    `json:"month"`
		Year       int    `json:"year"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	empID, err := primitive.ObjectIDFromHex(input.EmployeeID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	var emp models.Employee
	err = employeeCol.FindOne(context.Background(), bson.M{"_id": empID}).Decode(&emp)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Employee not found"})
		return
	}

	start := time.Date(input.Year, time.Month(input.Month), 1, 0, 0, 0, 0, time.UTC)
	end := start.AddDate(0, 1, 0)

	cursor, err := attendanceCol.Find(context.Background(), bson.M{
		"employee_id": empID,
		"date": bson.M{"$gte": start, "$lt": end},
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance"})
		return
	}
	defer cursor.Close(context.Background())

	totalHours := 0.0
	for cursor.Next(context.Background()) {
		var att models.Attendance
		if err := cursor.Decode(&att); err == nil {
			totalHours += att.WorkHours
		}
	}

	totalPay := emp.Salary * totalHours
	payroll := models.Payroll{
		ID:          primitive.NewObjectID(),
		EmployeeID:  empID,
		Month:       input.Month,
		Year:        input.Year,
		TotalHours:  totalHours,
		TotalPay:    totalPay,
		GeneratedAt: time.Now(),
	}

	_, err = payrollCol.InsertOne(context.Background(), payroll)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate payroll"})
		return
	}

	c.JSON(http.StatusCreated, payroll)
}

func GetPayrollHistory(c *gin.Context) {
	empID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(empID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	cursor, err := payrollCol.Find(context.Background(), bson.M{"employee_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch payrolls"})
		return
	}
	defer cursor.Close(context.Background())

	var records []models.Payroll
	for cursor.Next(context.Background()) {
		var rec models.Payroll
		if err := cursor.Decode(&rec); err == nil {
			records = append(records, rec)
		}
	}

	c.JSON(http.StatusOK, records)
}
