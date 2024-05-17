// course_handler.go

package handlers

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"

	"github.com/9500073161/coursemanagement/common"
	"github.com/9500073161/coursemanagement/statemanager"
)

// CourseHandler handles CRUD operations for courses
type CourseHandler struct {
	StateManager *statemanager.StateManager
}

// NewCourseHandler initializes a new CourseHandler
func NewCourseHandler(sm *statemanager.StateManager) *CourseHandler {
	return &CourseHandler{StateManager: sm}
}

// Create handles creating a new course

func (ch *CourseHandler) HandlersCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
    switch r.Method{ 
	case http.MethodPost:
		ch.CreateCourse(w,r)
	case http.MethodGet:
		ch.GetAllCourse(w,r)
	// case http.MethodPut:
	// 	ch.Update(w,r)
	// case http.MethodDelete:
	// 	ch.Delete(w,r)
	default :
		w.WriteHeader(http.StatusMethodNotAllowed)
		return				
	}
	
}

func (ch *CourseHandler) CreateCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

	var course common.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the course data (e.g., save it to a database)
	fmt.Printf("Received new course: %+v\n", course)
	ch.StateManager.CreateCourseEntry(course)
	// Respond with a success message
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Course created successfully")
}

// GetAll handles retrieving all courses
func (ch *CourseHandler) GetAllCourse(w http.ResponseWriter, r *http.Request) {
	// Implementation of get all operation

	// Get all courses from the state manager

	
	courses, err := ch.StateManager.GetAllCourses()

	// Convert courses to JSON
	response, err := json.Marshal(courses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond with the list of courses
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

