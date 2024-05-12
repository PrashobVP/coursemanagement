// common/course_repository.go

package common

import (
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"gorm.io/gorm"
)

type courseRepository struct {
	DB *gorm.DB
}

func NewcourseRepository(db *gorm.DB) *courseRepository {
	return &courseRepository{DB: db}
}

func (r *courseRepository) Createcourse(course *mysqldbmodels.Course) error {
	return r.DB.Create(course).Error
}

func (r *courseRepository) GetcourseByID(id uint) (*mysqldbmodels.Course, error) {
	var course mysqldbmodels.Course
	err := r.DB.First(&course, id).Error
	if err != nil {
		return nil, err
	}
	return &course, nil
}

func (r *courseRepository) Updatecourse(course *mysqldbmodels.Course) error {
	return r.DB.Save(course).Error
}

func (r *courseRepository) Deletecourse(id uint) error {
	return r.DB.Delete(&mysqldbmodels.Course{}, id).Error
}
