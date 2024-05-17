package main

import (
	"encoding/json"
	"fmt"
	"html"
	"log"
	"net/http"

	"github.com/9500073161/coursemanagement/common"
	"github.com/9500073161/coursemanagement/statemanager"
)

var sm *statemanager.StateManager

func main() {
	fmt.Println("Successfully completed CRUD operation")

	sm, err := statemanager.InitStateManager()
	// c1 := common.Course{TeacherID: 123}
	if err != nil {
		return
	}
	// sm.CreateCourseEntry(c1)

	http.HandleFunc("/home", homeHandler)

	http.HandleFunc("/course", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))

		var course common.Course
		err := json.NewDecoder(r.Body).Decode(&course)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Do something with the course data (e.g., save it to a database)
		fmt.Printf("Received new course: %+v\n", course)
		sm.CreateCourseEntry(course)
		// Respond with a success message
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Course created successfully")

	})

	http.HandleFunc("/student", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		var student common.Student
		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Do something with the course data (e.g., save it to a database)
		fmt.Printf("Received new Student: %+v\n", student)
		sm.CreateStudentEntry(student)
		// Respond with a success message
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Student created successfully")

	})

	http.HandleFunc("/teacher", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		var teacher common.Teacher
		err := json.NewDecoder(r.Body).Decode(&teacher)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Do something with the course data (e.g., save it to a database)
		fmt.Printf("Received new Teacher: %+v\n", teacher)
		sm.CreateTeacherEntry(teacher)
		// Respond with a success message
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Teacher created successfully")

	})

	http.HandleFunc("/entrollment", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
		var entrollment common.Entrollment
		err := json.NewDecoder(r.Body).Decode(&entrollment)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Do something with the course data (e.g., save it to a database)
		fmt.Printf("Received new entrollment: %+v\n", entrollment)
		sm.CreateEntrollmentEntry(entrollment)
		// Respond with a success message
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "entrollment created successfully")

	})

	log.Fatal(http.ListenAndServe(":10000", nil))

}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("homeHandler call")

}
