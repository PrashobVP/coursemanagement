// storage/database.go

package mysqldbmodels

import (
	"errors"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	Conn *gorm.DB
}

func InitializeDatabase() (*DBClient,error) {

	dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    if dbUser == "" || dbPassword == "" {
        return nil, errors.New("DB_USER or DB_PASSWORD environment variable not set")
    }

	
	// Connect to MySQL database
	dsn := dbUser + ":" + dbPassword + "@tcp(host.docker.internal:3306)/coursemanagement?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
        return nil, err
    }
	client := &DBClient{Conn: DB}
	return client,err

	
}
