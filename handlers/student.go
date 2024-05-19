// Student_handler.go

package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/9500073161/coursemanagement/common"
	"github.com/9500073161/coursemanagement/statemanager"
)

// StudentHandler handles CRUD operations for courses
type StudentHandler struct {
	StateManager *statemanager.StateManager
}

type UpdateStudentRequest struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// NewCourseHandler initializes a new StudentHandler
func NewStudentHandler(sm *statemanager.StateManager) *StudentHandler {
	return &StudentHandler{StateManager: sm}
}

// Create handles creating a new Student

func (ch *StudentHandler) HandlersStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	switch r.Method {
	case http.MethodPost:
		ch.CreateStudent(w, r)
	case http.MethodGet:
		ch.GetAllStudent(w, r)
	case http.MethodPut:
		ch.StudentUpdate(w, r)
	// case http.MethodDelete:
	// 	ch.Delete(w,r)
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

}

func (ch *StudentHandler) CreateStudent(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	var student common.Student
	err := json.NewDecoder(r.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the course data (e.g., save it to a database)
	fmt.Printf("Received new student: %+v\n", student)
	ch.StateManager.CreateStudentEntry(student)
	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Student created successfully")
}

// GetAll handles retrieving all courses
func (ch *StudentHandler) GetAllStudent(w http.ResponseWriter, r *http.Request) {
	// Implementation of get all operation

	// Get all courses from the state manager

	students, err := ch.StateManager.GetAllStudents()

	// Convert courses to JSON
	response, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of courses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	fmt.Fprintf(w, "Students Get successfully")
}

// Update handles updating an existing course
func (ch *StudentHandler) StudentUpdate(w http.ResponseWriter, r *http.Request) {

	var req UpdateStudentRequest
	// Decode the JSON request body into a course object
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

	// Perform the update operation using the StateManager
	students, err := ch.StateManager.UpdateAllStudents(req.ID, req.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response, err := json.Marshal(students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
	fmt.Fprintf(w, "Student updated successfully")

}

// // Delete handles deleting a course
// func (ch *CourseHandler) Delete(w http.ResponseWriter, r *http.Request) {
// 	// Implementation of delete operation

// 	var course common.Course
// 	err := json.NewDecoder(r.Body).Decode(&course)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// // Do something with the course data (e.g., delete it from the database)
// fmt.Printf("Delete course: %+v\n", course)
// ch.StateManager.DeleteCourse(course.ID) // Assuming course has an ID field
// // Respond with a success message
// w.WriteHeader(http.StatusOK)
// fmt.Fprintf(w, "Course deleted successfully")
