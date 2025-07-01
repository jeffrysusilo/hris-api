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
)

var attendanceCol = config.GetCollection("attendances")

func CheckIn(c *gin.Context) {
	var input struct {
		EmployeeID string `json:"employee_id"`
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

	today := time.Now().Truncate(24 * time.Hour)
	filter := bson.M{"employee_id": empID, "date": today}
	existing := attendanceCol.FindOne(context.Background(), filter)
	if existing.Err() == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already checked in today"})
		return
	}

	attendance := models.Attendance{
		ID:         primitive.NewObjectID(),
		EmployeeID: empID,
		CheckIn:    time.Now(),
		Date:       today,
		WorkHours:  0,
	}

	_, err = attendanceCol.InsertOne(context.Background(), attendance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check in"})
		return
	}

	c.JSON(http.StatusCreated, attendance)
}

func CheckOut(c *gin.Context) {
	attendanceID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(attendanceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid attendance ID"})
		return
	}

	var attendance models.Attendance
	err = attendanceCol.FindOne(context.Background(), bson.M{"_id": objID}).Decode(&attendance)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Attendance record not found"})
		return
	}

	if attendance.CheckOut != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Already checked out"})
		return
	}

	now := time.Now()
	hours := now.Sub(attendance.CheckIn).Hours()
	attendance.CheckOut = &now
	attendance.WorkHours = hours

	_, err = attendanceCol.UpdateOne(
		context.Background(),
		bson.M{"_id": objID},
		bson.M{"$set": bson.M{"check_out": now, "work_hours": hours}},
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check out"})
		return
	}

	c.JSON(http.StatusOK, attendance)
}

func GetEmployeeAttendance(c *gin.Context) {
	empID := c.Param("id")
	objID, err := primitive.ObjectIDFromHex(empID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid employee ID"})
		return
	}

	cursor, err := attendanceCol.Find(context.Background(), bson.M{"employee_id": objID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch attendance records"})
		return
	}
	defer cursor.Close(context.Background())

	var records []models.Attendance
	for cursor.Next(context.Background()) {
		var rec models.Attendance
		if err := cursor.Decode(&rec); err == nil {
			records = append(records, rec)
		}
	}

	c.JSON(http.StatusOK, records)
}
