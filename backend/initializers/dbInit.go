package initializers

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	fmt.Println("::: LOADING DATABASES :::")
	dbstr := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open(dbstr), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
