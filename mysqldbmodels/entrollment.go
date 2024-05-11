// mysqldbmodels/entrollment.go

package mysqldbmodels

import "fmt"

type EntrollmentRepository interface {
	Createentrollment(entrollment *Entrollment) error
	GetentrollmentByID(id uint) (*Entrollment, error)
	// Add other methods as needed
}

type Entrollment struct {
	ID   uint   `gorm:"primaryKey"`
	CourseID int
	StudentID int
}

func MainEntrollment(repo EntrollmentRepository) {
	// Create a new entrollment
	entrollment := &Entrollment{}
	if err := repo.Createentrollment(entrollment); err != nil {
		panic("failed to create entrollment")
	}
	fmt.Println("entrollment created successfully")

	// Read teacher
	retrievedentrollment, err := repo.GetentrollmentByID(entrollment.ID)
	if err != nil {
		panic("failed to retrieve entrollment")
	}
	fmt.Printf("Retrieved entrollment: %+v\n", *retrievedentrollment)

	// Add other CRUD operations
}
