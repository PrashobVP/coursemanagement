// mysqldbmodels/course.go

package mysqldbmodels

import "fmt"

type Course struct {
	ID        int    `gorm:"column:id;primaryKey"`
	TeacherID int    `gorm:"column:teacher_id;"`
	Name      string `gorm:"column:name;"`
}

func (client *DBClient) CreateCourseRow() error {
	fmt.Println("Started Course table creation")

	var c = Course{ID: 1, Name: "golang", TeacherID: 1}

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
