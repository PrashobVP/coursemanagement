// mysqldbmodels/teacher.go

package mysqldbmodels

import "fmt"

type Teacher struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;"`
}

func (client *DBClient) CreateTeacherRow() error {
	fmt.Println("Started Teacher table creation")

	var t = Teacher{ID: 1, Name: "Shubam"}

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
