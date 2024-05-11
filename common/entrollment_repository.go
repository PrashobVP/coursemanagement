// common/entrollment_repository.go

package common

import (
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"gorm.io/gorm"
)

type entrollmentRepository struct {
	DB *gorm.DB
}

// Createcourse implements mysqldbmodels.CourseRepository.
func (r *entrollmentRepository) Createcourse(course *mysqldbmodels.Course) error {
	panic("unimplemented")
}

// GetcourseByID implements mysqldbmodels.CourseRepository.
func (r *entrollmentRepository) GetcourseByID(id uint) (*mysqldbmodels.Course, error) {
	panic("unimplemented")
}

func NewentrollmentRepository(db *gorm.DB) *entrollmentRepository {
	return &entrollmentRepository{DB: db}
}

func (r *entrollmentRepository) Createentrollment(entrollment *mysqldbmodels.Entrollment) error {
	return r.DB.Create(entrollment).Error
}

func (r *entrollmentRepository) GetentrollmentByID(id uint) (*mysqldbmodels.Entrollment, error) {
	var entrollment mysqldbmodels.Entrollment
	err := r.DB.First(&entrollment, id).Error
	if err != nil {
		return nil, err
	}
	return &entrollment, nil
}

// Implement other CRUD operations as needed
