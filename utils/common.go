package utils

import (
	"github.com/go-playground/validator/v10"
	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Globals struct {
	App         *fiber.App
	DB          *gorm.DB
	Logger      *zap.Logger
	Validate    *validator.Validate
	RestyClient *resty.Client
}
