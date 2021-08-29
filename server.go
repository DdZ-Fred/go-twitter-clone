package gotwitterclone

import (
	"log"

	"github.com/DdZ-Fred/go-twitter-clone/api"
	"github.com/DdZ-Fred/go-twitter-clone/database"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/DdZ-Fred/go-twitter-clone/validation"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
)

func Run() {
	db := database.InitDB()               // DB Init
	logger, _ := zap.NewProduction()      // Zap Logger Init
	app := fiber.New()                    // Fiber App Init
	validate := validation.InitValidate() // Validate Tool Init

	globals := utils.Globals{
		App:      app,
		DB:       db,
		Logger:   logger,
		Validate: validate,
	}

	api.Api(globals)

	log.Fatal(app.Listen(":3001"))
}
