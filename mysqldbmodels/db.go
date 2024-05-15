// storage/database.go

package mysqldbmodels

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	Conn *gorm.DB
}

func InitializeDatabase() (*DBClient,error) {
	// Connect to MySQL database
	dsn := "root:1234@tcp(localhost:3306)/coursemanagement?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	client := &DBClient{Conn: DB}
	return client,err

	// // Migrate the schema
	// if err := DB.AutoMigrate(&mysqldbmodels.Teacher{}, &mysqldbmodels.Student{}, &mysqldbmodels.Course{}, &mysqldbmodels.Entrollment{}); err != nil {
	// 	panic("failed to migrate schema")
	// }
	

	//return client
}
