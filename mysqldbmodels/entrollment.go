package mysqldbmodels

import "fmt"

type EntrollmentRepository interface {
    Createentrollment(entrollment *Entrollment) error
    GetentrollmentByID(id uint) (*Entrollment, error)
    Updateentrollment(entrollment *Entrollment) error
    Deleteentrollment(id uint) error
    // Add other methods as needed
}

type Entrollment struct {
    ID        uint `gorm:"primaryKey"`
    CourseID  int
    StudentID int
}

func MainEntrollment(repo EntrollmentRepository) {
    // Create a new entrollment
    entrollment := &Entrollment{}
    if err := repo.Createentrollment(entrollment); err != nil {
        panic("failed to create entrollment")
    }
    fmt.Println("entrollment created successfully")

    // Read entrollment
    retrievedentrollment, err := repo.GetentrollmentByID(entrollment.ID)
    if err != nil {
        panic("failed to retrieve entrollment")
    }
    fmt.Printf("Retrieved entrollment: %+v\n", *retrievedentrollment)

    // Update entrollment
    retrievedentrollment.CourseID = 123 // Example update
    if err := repo.Updateentrollment(retrievedentrollment); err != nil {
        panic("failed to update entrollment")
    }
    fmt.Println("entrollment updated successfully")

    // Delete entrollment
    if err := repo.Deleteentrollment(entrollment.ID); err != nil {
        panic("failed to delete entrollment")
    }
    fmt.Println("entrollment deleted successfully")
}
