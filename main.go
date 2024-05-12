package main

import (
	"fmt"

	"github.com/9500073161/coursemanagement/mysqldbmodels"
)

func main() {
	fmt.Println("Successfully completed CRUD operation")

	// Initialize the database
	db := mysqldbmodels.InitializeDatabase()
	db.CreateCourseRow()
	db.CreateStudentRow()
	db.CreateTeacherRow()
	db.CreateEntrollmentRow()
	

	

	
}
