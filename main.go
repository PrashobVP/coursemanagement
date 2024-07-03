package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/9500073161/coursemanagement/handlers"
	"github.com/9500073161/coursemanagement/statemanager"
)

func main() {
	// Initialize the state manager
	sm, err := statemanager.InitStateManager()
	if err != nil {
		log.Fatal("Failed to initialize state manager:", err)
	}

	// Create handlers
	courseHandler := handlers.NewCourseHandler(sm)
	// Similarly, create handlers for other entities
	studentHandler := handlers.NewStudentHandler(sm)
	teacherHandler := handlers.NewTeacherHandler(sm)
	enrollmentHandler := handlers.NewEntrollmentHandler(sm)

	// Define HTTP routes
	http.HandleFunc("/course", courseHandler.HandlersCourse)
	http.HandleFunc("/student", studentHandler.HandlersStudent)
	http.HandleFunc("/teacher", teacherHandler.HandlersTeacher)
	http.HandleFunc("/entrollment", enrollmentHandler.HandlersEntrollment)

	fmt.Println("Server is running...")
	log.Fatal(http.ListenAndServe(":5001", nil))
}
