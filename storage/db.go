// storage/database.go

package storage

import (
    "github.com/9500073161/coursemanagement/mysqldbmodels"
    "gorm.io/driver/mysql"
    "gorm.io/gorm"
)

var DB *gorm.DB

func InitializeDatabase() *gorm.DB {
    // Connect to MySQL database
    dsn := "root:1234@tcp(localhost:3306)/coursemanagement?charset=utf8mb4&parseTime=True&loc=Local"
    var err error
    DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        panic("failed to connect database")
    }

    // Migrate the schema
    if err := DB.AutoMigrate(&mysqldbmodels.Teacher{}, &mysqldbmodels.Student{}, &mysqldbmodels.Course{}, &mysqldbmodels.Entrollment{}); err != nil {
        panic("failed to migrate schema")
    }

    return DB
}
