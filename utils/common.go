package utils

import (
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/go-redis/redis/v8"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Middlewares struct {
	JwtAuth fiber.Handler
}

type Globals struct {
	Ctx         *context.Context
	App         *fiber.App
	DB          *gorm.DB
	RDB         *redis.Client
	Logger      *zap.Logger
	Validate    *validator.Validate
	RestyClient *resty.Client
	Middlewares *Middlewares
}
