package controllers

import "hris-api/config"

func InitControllers() {
	EmployeeCol = config.GetCollection("employees")
	AttendanceCol = config.GetCollection("attendances")
	PayrollCol = config.GetCollection("payrolls")
}
