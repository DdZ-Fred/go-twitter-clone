package gotwitterclone

import (
	"log"

	"github.com/DdZ-Fred/go-twitter-clone/api"
	"github.com/DdZ-Fred/go-twitter-clone/db_postgres"
	"github.com/DdZ-Fred/go-twitter-clone/db_redis"
	"github.com/DdZ-Fred/go-twitter-clone/emails"
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

	// ctx := context.Background()                  // !! CONTEXT MUST NOT BE STOTED IN STRUCTS Redis dependencies for all operations
	db := db_postgres.InitDB()                   // DB Init
	rdb := db_redis.InitDB()                     // Redis DB Init
	logger, _ := zap.NewProduction()             // Zap Logger Init
	validate := validation.InitValidate()        // Validate Tool Init
	restyClient := resty.New()                   // Resty HTTP Client
	middlewares := middlewares.InitMiddlewares() // Fiber middlewares
	emails := emails.InitEmails()

	globals := utils.Globals{
		// Ctx:         &ctx,
		App:         app,
		DB:          db,
		RDB:         rdb,
		Logger:      logger,
		Validate:    validate,
		RestyClient: restyClient,
		Middlewares: middlewares,
		Emails:      emails,
	}

	api.Api(globals)

	log.Fatal(app.Listen(":3001"))
}
