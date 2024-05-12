// mysqldbmodels/student.go

package mysqldbmodels

import "fmt"

type Student struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;"`
}

func (client *DBClient) CreateStudentRow() error {
	fmt.Println("Started Student table creation")

	var s = Student{ID:1, Name: "Kanchi"}

	var err error

	db := client.Conn
	tx := db.Begin()

	if err = tx.Create(&s).Error; err != nil {
		//utils.Logger.Error("error while creating a position row from db", zap.Error(err))
		tx.Rollback()
		return err
	}
	tx.Commit()
	//utils.Logger.Warn("successfully created a position row in db")
	return err
}
