package initializers

import "github.com/kiridharan/models"

func AutoSync() {

	DB.AutoMigrate(&models.User{})
}
