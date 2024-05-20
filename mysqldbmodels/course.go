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

//getallcourses

func (client *DBClient) GetCourseRaw() ([]Course, error) {
	var courses []Course
	fmt.Println("Get all course from DB")
	query := `select 
			  *
			from courses`

	db := client.Conn

	if err := db.Raw(query).Find(&courses).Error; err != nil {
		return courses, fmt.Errorf("error while reading rows from positions table,err: %s", err.Error())
	}
	return courses, nil

}

//update course

func (client *DBClient) UpdateCourseRaw(ID int, Name string, TeacherID int) ([]Course, error) {
	var courses []Course
	fmt.Println("Updating course in DB")
	query := `UPDATE courses
	          SET name = ?, teacher_id = ?
	          WHERE id = ?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query, Name, TeacherID, ID).Error; err != nil {
		return courses, fmt.Errorf("error while updating course: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&courses).Error; err != nil {
		return courses, fmt.Errorf("error fetching updated course: %s", err.Error())
	}

	return courses, nil
}

//delete course

func (client *DBClient) DeleteCourseRaw(ID int) ([]Course, error){

	var courses []Course
	fmt.Println("Deleting course in DB")
	query := `DELETE FROM courses WHERE id=?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query,ID).Error; err != nil {
		return courses, fmt.Errorf("error while updating course: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&courses).Error; err != nil {
		return courses, fmt.Errorf("error fetching updated course: %s", err.Error())
	}

	return courses, nil

}


