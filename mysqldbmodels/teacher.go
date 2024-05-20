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

	var t = Teacher{ID: t1.ID, Name: t1.Name}

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

//getallteacher

func (client *DBClient) GetTeacherRaw() ([]Teacher, error) {
	var teachers []Teacher

	fmt.Println("Get all teachers from DB")
	query := `select 
			  *
			from teachers`

	db := client.Conn

	if err := db.Raw(query).Find(&teachers).Error; err != nil {
		return teachers, fmt.Errorf("error while reading rows from positions table,err: %s", err.Error())
	}
	return teachers, nil
}

//update Teacher

func (client *DBClient) UpdateTeacherRaw(ID int, Name string) ([]Teacher, error) {
	var teachers []Teacher
	fmt.Println("Updating teacher in DB")
	query := `UPDATE teachers
	          SET name = ?
	          WHERE id = ?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query, Name, ID).Error; err != nil {
		return teachers, fmt.Errorf("error while updating teacher: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&teachers).Error; err != nil {
		return teachers, fmt.Errorf("error fetching updated teacher: %s", err.Error())
	}

	return teachers, nil
}

//delete teacher

func (client *DBClient) DeleteTeacherRaw(ID int) ([]Teacher, error){

	var teachers []Teacher
	fmt.Println("Deleting teachers in DB")
	query := `DELETE FROM teachers WHERE id=?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query,ID).Error; err != nil {
		return teachers, fmt.Errorf("error while updating teachers: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&teachers).Error; err != nil {
		return teachers, fmt.Errorf("error fetching updated teachers: %s", err.Error())
	}

	return teachers, nil

}

