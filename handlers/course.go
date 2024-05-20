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

type UpdateCourseRequest struct {
	ID int  `json:"id"`
	Name      string `json:"name"`
	TeacherID int    `json:"teacher_id"`
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
	case http.MethodPut:
	    ch.CourseUpdate(w,r)
	case http.MethodDelete:
	    ch.DeleteCourseIDs(w,r)
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

// Update handles updating an existing course
func (ch *CourseHandler) CourseUpdate(w http.ResponseWriter, r *http.Request) {

	var req UpdateCourseRequest
    // Decode the JSON request body into a course object
    if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, fmt.Sprintf("Invalid request body: %s", err.Error()), http.StatusBadRequest)
		return
	}

    // Perform the update operation using the StateManager
    courses, err := ch.StateManager.UpdateAllCourses(req.ID, req.Name,req.TeacherID)
	if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
	response, err := json.Marshal(courses)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Respond with a success message
	w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
	w.Write(response)
    
}




//Delete handles deleting a course
func (ch *CourseHandler) DeleteCourseIDs(w http.ResponseWriter, r *http.Request) {


   var course common.Course
	err := json.NewDecoder(r.Body).Decode(&course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Do something with the course data (e.g., delete it from the database)
	fmt.Printf("Delete course: %+v\n", course)
	ch.StateManager.DeleteCourseID(course.ID) // Assuming course has an ID field
	// Respond with a success message
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Course deleted successfully")

}