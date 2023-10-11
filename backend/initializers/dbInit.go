package initializers

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func DatabaseInit() {

	fmt.Println("::: LOADING DATABASES :::")
	// dbstr := os.Getenv("DB")
	var err error
	DB, err = gorm.Open(postgres.Open("host=localhost user=postgres password=123456789 dbname=golang port=5432"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

}
