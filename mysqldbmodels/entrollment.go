// mysqldbmodels/entrollment.go

package mysqldbmodels

import "fmt"

type Entrollment struct {
	ID   int    `gorm:"column:id;primaryKey"`
	CourseID int `gorm:"column:course_id;"`
	StudentID int `gorm:"column:student_id;"`
}

func (client *DBClient) CreateEntrollmentRow() error {
	fmt.Println("Started Entrollment table creation")

	var e = Entrollment{ID: 1, CourseID: 1, StudentID: 1}

	var err error

	db := client.Conn
	tx := db.Begin()

	if err = tx.Create(&e).Error; err != nil {
		//utils.Logger.Error("error while creating a position row from db", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	//utils.Logger.Warn("successfully created a position row in db")
	return err
}
