package db_postgres

import (
	"github.com/DdZ-Fred/go-twitter-clone/models"
	"gorm.io/gorm"
)

func createOrMigrate(db *gorm.DB) {
	// Required by User model:
	// https://github.com/go-gorm/gorm/issues/1978#issuecomment-476673540
	db.Exec("CREATE TYPE user_status AS ENUM ('pending', 'active')")

	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Tweet{})
}
