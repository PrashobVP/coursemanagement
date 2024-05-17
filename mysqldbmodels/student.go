// mysqldbmodels/student.go

package mysqldbmodels

import (
	"fmt"

	"github.com/9500073161/coursemanagement/common"
)

type Student struct {
	ID   int    `gorm:"column:id;primaryKey"`
	Name string `gorm:"column:name;"`
}

func (client *DBClient) CreateStudentRow(s1 common.Student) error {
	fmt.Println("Started Student table creation")

	var s = Student{ID: s1.ID, Name: s1.Name}

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

//getallstudents

func (client *DBClient) GetStudentRaw() ([]Student,error) {
    var students []Student

	fmt.Println("Get all course from DB")
	query := `select 
			  *
			from students`
	
	db := client.Conn
	
	if err := db.Raw(query).Find(&students).Error; err != nil {
		return students, fmt.Errorf("error while reading rows from positions table,err: %s", err.Error())
	}
	return students, nil
}