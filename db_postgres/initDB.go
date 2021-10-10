package db_postgres

import (
	"fmt"
	"os"

	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	dsn := fmt.Sprintf(`
		host=%s
		user=%s
		password=%s
		dbname=%s
		port=%s
		sslmode=disable
		TimeZone=UTC`,
		os.Getenv("POSTGRES_HOST"),
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		os.Getenv("POSTGRES_DB"),
		os.Getenv("POSTGRES_PORT"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	runArgs := os.Args

	if utils.Contains(runArgs, "-cm") {
		createOrMigrate(db)
	}

	return db
}
