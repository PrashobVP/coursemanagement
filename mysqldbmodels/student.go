package mysqldbmodels

import "fmt"

type StudentRepository interface {
    CreateStudent(student *Student) error
    GetStudentByID(id uint) (*Student, error)
    UpdateStudent(student *Student) error
    DeleteStudent(id uint) error
    // Add other methods as needed
}

type Student struct {
    ID   uint   `gorm:"primaryKey"`
    Name string
}

func MainStudent(repo StudentRepository) {
    // Create a new student
    student := &Student{Name: "Prashob"}
    if err := repo.CreateStudent(student); err != nil {
        panic("failed to create student")
    }
    fmt.Println("Student created successfully")

    // Read student
    retrievedStudent, err := repo.GetStudentByID(student.ID)
    if err != nil {
        panic("failed to retrieve student")
    }
    fmt.Printf("Retrieved Student: %+v\n", *retrievedStudent)

    // Update student
    retrievedStudent.Name = "Prashob Updated"
    if err := repo.UpdateStudent(retrievedStudent); err != nil {
        panic("failed to update student")
    }
    fmt.Println("Student updated successfully")

    // Delete student
    if err := repo.DeleteStudent(student.ID); err != nil {
        panic("failed to delete student")
    }
    fmt.Println("Student deleted successfully")
}
