// common/student_repository.go

package common

import (
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func NewStudentRepository(db *gorm.DB) *StudentRepository {
	return &StudentRepository{DB: db}
}

func (r *StudentRepository) CreateStudent(student *mysqldbmodels.Student) error {
	return r.DB.Create(student).Error
}

func (r *StudentRepository) GetStudentByID(id uint) (*mysqldbmodels.Student, error) {
	var student mysqldbmodels.Student
	err := r.DB.First(&student, id).Error
	if err != nil {
		return nil, err
	}
	return &student, nil
}

func (r *StudentRepository) UpdateStudent(student *mysqldbmodels.Student) error {
	return r.DB.Save(student).Error
}

func (r *StudentRepository) DeleteStudent(id uint) error {
	return r.DB.Delete(&mysqldbmodels.Student{}, id).Error
}
