// common/entrollment_repository.go

package common

import (
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"gorm.io/gorm"
)

type entrollmentRepository struct {
	DB *gorm.DB
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

func (r *entrollmentRepository) Updateentrollment(entrollment *mysqldbmodels.Entrollment) error {
	return r.DB.Save(entrollment).Error
}

func (r *entrollmentRepository) Deleteentrollment(id uint) error {
	return r.DB.Delete(&mysqldbmodels.Entrollment{}, id).Error
}
