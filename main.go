// main.go

package main

import (
	"github.com/9500073161/coursemanagement/common"
	"github.com/9500073161/coursemanagement/mysqldbmodels"
	"github.com/9500073161/coursemanagement/storage"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func init() {
	storage.InitializeDatabase()
}

func main() {

	// Connect to MySQL database
	dsn := "root:1234@tcp(localhost:3306)/coursemanagement?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Initialize repository
	teacherRepo := common.NewTeacherRepository(db)
	studentRepo := common.NewStudentRepository(db)
	courseRepo := common.NewcourseRepository(db)
	entrollmentRepo := common.NewentrollmentRepository(db)

	// Call Mainsubmain function
	mysqldbmodels.MainTeacher(teacherRepo)
	mysqldbmodels.MainStudent(studentRepo)
	mysqldbmodels.MainCourse(courseRepo)
	mysqldbmodels.MainEntrollment(entrollmentRepo)

}
