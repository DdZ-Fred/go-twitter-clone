package handlers

import (
	"errors"
	"time"

	apierrors "github.com/DdZ-Fred/go-twitter-clone/api-errors"
	"github.com/DdZ-Fred/go-twitter-clone/jwt"
	"github.com/DdZ-Fred/go-twitter-clone/models"
	"github.com/DdZ-Fred/go-twitter-clone/password"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/DdZ-Fred/go-twitter-clone/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Auth struct {
	Globals utils.Globals
}

func (auth Auth) Login() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var input LoginCredentials
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				apierrors.ErrorResponseBadPayloadFormat(
					err.Error(),
				),
			)
		}
		if len(input.Username) == 0 || len(input.Password) == 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		var user models.User

		if err := auth.Globals.DB.Model(&models.User{}).First(&user, "username = ?", input.Username).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
		}

		match, err := password.ComparePasswordAndHash(
			input.Password,
			user.Password,
		)
		if err != nil {
			auth.Globals.Logger.Fatal(
				"Couldn't compare password with encoded hash",
				zap.Error(err),
			)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		if !match {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		tokenString, expirationTime, err := jwt.GetJWT(auth.Globals, user)

		if err != nil {
			auth.Globals.Logger.Fatal(
				"Couldn't generate a JWT",
				zap.Error(err),
			)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		c.Cookie(&fiber.Cookie{
			Name:     "token",
			Value:    tokenString,
			Expires:  expirationTime,
			Secure:   true,
			HTTPOnly: true,
		})

		return c.JSON(fiber.Map{
			"accessToken": tokenString,
		})
	}
}

type SignupPayload struct {
	Fname     string `validate:"required,min=2,max=30" json:"fname"`
	Lname     string `validate:"required,min=2,max=30" json:"lname"`
	Email     string `validate:"required,email" json:"email"`
	Username  string `validate:"required,min=3,max=30" json:"username"`
	BirthDate string `validate:"required,datetime=2006-01-02T15:04:05Z07:00" json:"birthDate"`
	Country   string `validate:"required,iso3166_1_alpha2" json:"country"`
	Password  string `validate:"required,password" json:"password"`
}

func (auth Auth) SignUp() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload SignupPayload

		// Payload parsing
		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				apierrors.ErrorResponseBadPayloadFormat(
					err.Error(),
				),
			)
		}

		// Struct validation
		validationErrors := validation.ExtractValidationErrorsFromErr(
			auth.Globals.Validate.Struct(payload),
		)
		if validationErrors != nil {
			return validation.FailedValidationResponse(c, validationErrors)
		}

		// Model init
		birthDate, _ := time.Parse(time.RFC3339, payload.BirthDate)
		password, _ := password.GenerateHashFromPassword(payload.Password, &password.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		})

		newUser := models.User{
			Id:        uuid.New().String(),
			Fname:     payload.Fname,
			Lname:     payload.Lname,
			Email:     payload.Email,
			Username:  payload.Username,
			BirthDate: birthDate,
			Country:   payload.Country,
			Password:  password,
		}

		if err := auth.Globals.DB.Create(&newUser).Error; err != nil {
			// Read about error value: https://blog.golang.org/go1.13-errors#TOC_2.1.
			if pgErr, ok := err.(*pgconn.PgError); ok {
				switch pgErr.Code {
				// Unique constraint violation
				case "23505":
					var apiErrorCodeStatus apierrors.ApiErrorCodeStatus
					if pgErr.ConstraintName == "users_email_key" {
						apiErrorCodeStatus = apierrors.Auth["email_already_taken"]
					} else {
						apiErrorCodeStatus = apierrors.Auth["username_already_taken"]
					}
					return c.Status(fiber.StatusConflict).JSON(
						apierrors.ErrorResponseDataConflict(apiErrorCodeStatus),
					)
				}
			}
			auth.Globals.Logger.Fatal(
				"Error while creating a new user",
				zap.Error(err),
			)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.Status(201).JSON(newUser.ToUserSafe())
	}
}
