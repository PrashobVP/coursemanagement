package main

import (
	"fmt"

	"github.com/9500073161/coursemanagement/common"
	"github.com/9500073161/coursemanagement/handler"
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"github.com/9500073161/coursemanagement/storage"
	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("Successfully completed CRUD operation")

	// Initialize the database
	db := storage.InitializeDatabase()

	// Initialize repositories
	teacherRepo := common.NewTeacherRepository(db)
	studentRepo := common.NewStudentRepository(db)
	courseRepo := common.NewcourseRepository(db)
	entrollmentRepo := common.NewentrollmentRepository(db)

	// Call Main functions
	mysqldbmodels.MainTeacher(teacherRepo)
	mysqldbmodels.MainStudent(studentRepo)
	mysqldbmodels.MainCourse(courseRepo)
	mysqldbmodels.MainEntrollment(entrollmentRepo)

	// Initialize the course handler with the course repository
	courseHandler := handler.NewCourseHandler(courseRepo)

	// Initialize Gin router
	router := gin.Default()

	// Define routes
	router.POST("/courses", courseHandler.CreateCourse)
	router.GET("/courses/:id", courseHandler.GetCourseByID)
	router.PUT("/courses/:id", courseHandler.UpdateCourse)
	router.DELETE("/courses/:id", courseHandler.DeleteCourse)

	// Start the server
	router.Run(":8080")
}
