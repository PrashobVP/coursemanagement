// main.go

package main

import (
	"fmt"

	"github.com/9500073161/coursemanagement/storage"
)

func init() {
	storage.InitializeDatabase()
}

func main() {

	fmt.Println("Successfully completed CRUD operation")

	

}
