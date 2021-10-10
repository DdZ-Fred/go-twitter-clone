package db_postgres

import (
	"github.com/DdZ-Fred/go-twitter-clone/models"
	"gorm.io/gorm"
)

func createOrMigrate(db *gorm.DB) {
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Tweet{})
}
