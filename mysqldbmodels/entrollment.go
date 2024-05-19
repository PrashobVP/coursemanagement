// mysqldbmodels/entrollment.go

package mysqldbmodels

import (
	"fmt"

	"github.com/9500073161/coursemanagement/common"
)

type Entrollment struct {
	ID   int    `gorm:"column:id;primaryKey"`
	CourseID int `gorm:"column:course_id;"`
	StudentID int `gorm:"column:student_id;"`
}

func (client *DBClient) CreateEntrollmentRow(e1 common.Entrollment) error {
	fmt.Println("Started Entrollment table creation")

	var e = Entrollment{ID: e1.ID, CourseID: e1.CourseID, StudentID: e1.StudentID}

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

//getallentrollment

func (client *DBClient) GetEntrollmentRaw() ([]Entrollment,error) {
    var entrollments []Entrollment

	fmt.Println("Get all teachers from DB")
	query := `select 
			  *
			from entrollments`
	
	db := client.Conn
	
	if err := db.Raw(query).Find(&entrollments).Error; err != nil {
		return entrollments, fmt.Errorf("error while reading rows from positions table,err: %s", err.Error())
	}
	return entrollments, nil
}


//update Entrollment

func (client *DBClient) UpdateEntrollmentRaw(ID int, CourseID int, StudentID int) ([]Entrollment, error) {
	var entrollments []Entrollment
	fmt.Println("Updating Entrollment in DB")
	query := `UPDATE entrollments
	          SET course_id = ?, student_id = ?
	          WHERE id = ?`

	db := client.Conn

	// Execute the raw SQL query with placeholders
	if err := db.Exec(query, CourseID,StudentID, ID).Error; err != nil {
		return entrollments, fmt.Errorf("error while updating teacher: %s", err.Error())
	}

	// Fetch the updated course to return it
	if err := db.Where("id = ?", ID).First(&entrollments).Error; err != nil {
		return entrollments, fmt.Errorf("error fetching updated teacher: %s", err.Error())
	}

	return entrollments, nil
}

