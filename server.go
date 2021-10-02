package gotwitterclone

import (
	"log"

	"github.com/DdZ-Fred/go-twitter-clone/api"
	"github.com/DdZ-Fred/go-twitter-clone/database"
	"github.com/DdZ-Fred/go-twitter-clone/middlewares"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/DdZ-Fred/go-twitter-clone/validation"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

func Run() {
	app := fiber.New()            // Fiber App Init
	app.Use(cors.New(cors.Config{ // Check https://docs.gofiber.io/api/middleware/cors for more options
		AllowOrigins: "http://localhost:3000",
	}))

	db := database.InitDB()               // DB Init
	logger, _ := zap.NewProduction()      // Zap Logger Init
	validate := validation.InitValidate() // Validate Tool Init
	restyClient := resty.New()            // Resty HTTP Client
  middlewares := middlewares.InitMiddlewares()

	globals := utils.Globals{
		App:         app,
		DB:          db,
		Logger:      logger,
		Validate:    validate,
		RestyClient: restyClient,
    Middlewares: middlewares,
	}

	api.Api(globals)

	log.Fatal(app.Listen(":3001"))
}
