// storage/database.go

package mysqldbmodels

import (
	"errors"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBClient struct {
	Conn *gorm.DB
}

func InitializeDatabase() (*DBClient, error) {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	if dbUser == "" || dbPassword == "" {
		return nil, errors.New("DB_USER or DB_PASSWORD environment variable not set")
	}

	// Connect to MySQL database
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbHost, dbPort, dbName)
	var err error
	DB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	client := &DBClient{Conn: DB}
	return client, err

}
