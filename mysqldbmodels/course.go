// mysqldbmodels/course.go

package mysqldbmodels

import "fmt"

type CourseRepository interface {
	Createcourse(course *Course) error
	GetcourseByID(id uint) (*Course, error)
	// Add other methods as needed
}

type Course struct {
	ID   uint   `gorm:"primaryKey"`
	TeacherID int
	Name string
}

func MainCourse(repo CourseRepository) {
	// Create a new course
	course := &Course{Name: "Golang"}
	if err := repo.Createcourse(course); err != nil {
		panic("failed to create studnet")
	}
	fmt.Println("course created successfully")

	// Read course
	retrievedCourse, err := repo.GetcourseByID(course.ID)
	if err != nil {
		panic("failed to retrieve course")
	}
	fmt.Printf("Retrieved course: %+v\n", *retrievedCourse)

	// Add other CRUD operations
}
