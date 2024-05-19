// Entrollment_handler.go

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
type EntrollmentHandler struct {
	StateManager *statemanager.StateManager
}

type UpdateEntrollmentRequest struct {
	ID int  `json:"id"`
	CourseID      int `json:"course_id"`
	StudentID int    `json:"student_id"`
}

// NewCourseHandler initializes a new StudentHandler
func NewEntrollmentHandler(sm *statemanager.StateManager) *EntrollmentHandler {
	return &EntrollmentHandler{StateManager: sm}
}

// Create handles creating a new Student

func (ch *EntrollmentHandler) HandlersEntrollment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    switch r.Method{ 
	case http.MethodPost:
		ch.CreateEntrollment(w,r)
	case http.MethodGet:
		ch.GetAllEntrollment(w,r)
	case http.MethodPut:
	 	ch.EntrollmentUpdate(w,r)
	// case http.MethodDelete:
	// 	ch.Delete(w,r)
	default :
		w.WriteHeader(http.StatusMethodNotAllowed)
		return				
	}
	
}

func (ch *EntrollmentHandler) CreateEntrollment(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	var entrollment common.Entrollment
	err := json.NewDecoder(r.Body).Decode(&entrollment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the course data (e.g., save it to a database)
	fmt.Printf("Received new entrollment: %+v\n", entrollment)
	ch.StateManager.CreateEntrollmentEntry(entrollment)
	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "entrollment created successfully")
}

// GetAll handles retrieving all courses
func (ch *EntrollmentHandler) GetAllEntrollment(w http.ResponseWriter, r *http.Request) {
	// Implementation of get all operation

	// Get all courses from the state manager

	
	entrollments, err := ch.StateManager.GetAllEntrollments()

	// Convert courses to JSON
	response, err := json.Marshal(entrollments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of courses
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}


// Update handles updating an existing course
func (ch *EntrollmentHandler) EntrollmentUpdate(w http.ResponseWriter, r *http.Request) {

	var req UpdateEntrollmentRequest
    // Decode the JSON request body into a course object
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

    // Perform the update operation using the StateManager
    entrollments, err := ch.StateManager.UpdateAllEntrollments(req.ID, req.CourseID,req.StudentID)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	response, err := json.Marshal(entrollments)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with a success message
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(response)
    
}

// // Update handles updating an existing course
// func (ch *CourseHandler) Update(w http.ResponseWriter, r *http.Request) {
// 	// Implementation of update operation

// 	var course common.Course
// 	err := json.NewDecoder(r.Body).Decode(&course)
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		return
// 	}

// 	// Do something with the course data (e.g., update it in the database)
// 	fmt.Printf("Update course: %+v\n", course)
// 	ch.StateManager.UpdateCourse(course)
// 	// Respond with a success message
// 	w.WriteHeader(http.StatusOK)
// 	fmt.Fprintf(w, "Course updated successfully")
// }

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

