package gotwitterclone

import (
	"log"

	"github.com/DdZ-Fred/go-twitter-clone/database"
	"github.com/gofiber/fiber/v2"
)

func Run() {
	// Db Init
	database.InitDB()

	// Fiber app init
	app := fiber.New()

	log.Fatal(app.Listen(":3000"))
}
