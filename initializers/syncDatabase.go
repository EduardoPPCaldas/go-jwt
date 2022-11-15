package initializers

import "github.com/EduardoPPCaldas/go-jwt/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}