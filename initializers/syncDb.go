package initializers

import "golang-rnd/models"

func SyncDb() {
	DB.AutoMigrate(&models.LoginRequest{})
}