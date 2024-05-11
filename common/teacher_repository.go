// common/teacher_repository.go

package common

import (
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"gorm.io/gorm"
)

type TeacherRepository struct {
	DB *gorm.DB
}

func NewTeacherRepository(db *gorm.DB) *TeacherRepository {
	return &TeacherRepository{DB: db}
}

func (r *TeacherRepository) CreateTeacher(teacher *mysqldbmodels.Teacher) error {
	return r.DB.Create(teacher).Error
}

func (r *TeacherRepository) GetTeacherByID(id uint) (*mysqldbmodels.Teacher, error) {
	var teacher mysqldbmodels.Teacher
	err := r.DB.First(&teacher, id).Error
	if err != nil {
		return nil, err
	}
	return &teacher, nil
}

// Implement other CRUD operations as needed
