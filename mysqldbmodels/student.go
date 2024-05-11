// mysqldbmodels/student.go

package mysqldbmodels

import "fmt"

type StudentRepository interface {
	CreateStudent(student *Student) error
	GetStudentByID(id uint) (*Student, error)
	// Add other methods as needed
}

type Student struct {
	ID   uint   `gorm:"primaryKey"`
	Name string
}

func MainStudent(repo StudentRepository) {
	// Create a new Student
	student := &Student{Name: "Prashob"}
	if err := repo.CreateStudent(student); err != nil {
		panic("failed to create studnet")
	}
	fmt.Println("student created successfully")

	// Read teacher
	retrievedStudent, err := repo.GetStudentByID(student.ID)
	if err != nil {
		panic("failed to retrieve student")
	}
	fmt.Printf("Retrieved Student: %+v\n", *retrievedStudent)

	// Add other CRUD operations
}
