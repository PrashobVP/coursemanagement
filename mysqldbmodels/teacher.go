// mysqldbmodels/teacher.go

package mysqldbmodels

import (
	"fmt"

	"github.com/9500073161/coursemanagement/common"
)

type Teacher struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;"`
}

func (client *DBClient) CreateTeacherRow(t1 common.Teacher) error {
	fmt.Println("Started Teacher table creation")

	var t = Teacher{ID: t1.ID,Name: t1.Name }

	var err error

	db := client.Conn
	tx := db.Begin()

	if err = tx.Create(&t).Error; err != nil {
		//utils.Logger.Error("error while creating a position row from db", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	//utils.Logger.Warn("successfully created a position row in db")
	return err
}
