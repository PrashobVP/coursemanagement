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

func (client *DBClient) GetStudentRaw() ([]Student, error) {
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

//update student

func (client *DBClient) UpdateStudentRaw(ID int, Name string) ([]Student, error) {
	var students []Student
	fmt.Println("Updating course in DB")
	query := `UPDATE students
	          SET name = ?
	          WHERE id = ?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query, Name, ID).Error; err != nil {
		return students, fmt.Errorf("error while updating student: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&students).Error; err != nil {
		return students, fmt.Errorf("error fetching updated student: %s", err.Error())
	}

	return students, nil
}

//delete student

func (client *DBClient) DeleteStudentRaw(ID int) ([]Student, error){

	var students []Student
	fmt.Println("Deleting Student in DB")
	query := `DELETE FROM students WHERE id=?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query,ID).Error; err != nil {
		return students, fmt.Errorf("error while updating course: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&students).Error; err != nil {
		return students, fmt.Errorf("error fetching updated course: %s", err.Error())
	}

	return students, nil

}
