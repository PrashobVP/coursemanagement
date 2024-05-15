// mysqldbmodels/course.go

package mysqldbmodels

import (
	"fmt"

	"github.com/9500073161/coursemanagement/common"
)

type Course struct {
	ID        int    `gorm:"column:id;primaryKey"`
	TeacherID int    `gorm:"column:teacher_id;"`
	Name      string `gorm:"column:name;"`
}

func (client *DBClient) CreateCourseRow(c1 common.Course) error {
	fmt.Println("Started Course table creation")

	var c = Course{ID: c1.ID, Name: c1.Name, TeacherID: c1.TeacherID}

	var err error

	db := client.Conn
	tx := db.Begin()

	if err = tx.Create(&c).Error; err != nil {
		//utils.Logger.Error("error while creating a position row from db", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	//utils.Logger.Warn("successfully created a position row in db")
	return err

}
