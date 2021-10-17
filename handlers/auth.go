package handlers

import (
	"errors"
	"strings"
	"time"

	"github.com/DdZ-Fred/go-twitter-clone/api_errors"
	"github.com/DdZ-Fred/go-twitter-clone/emails"
	gormtypes "github.com/DdZ-Fred/go-twitter-clone/gorm_types"
	"github.com/DdZ-Fred/go-twitter-clone/jwt"
	"github.com/DdZ-Fred/go-twitter-clone/models"
	"github.com/DdZ-Fred/go-twitter-clone/password"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/DdZ-Fred/go-twitter-clone/validation"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/jackc/pgconn"
	"go.uber.org/zap"
	"gopkg.in/guregu/null.v4"
	"gorm.io/gorm"
)

type LoginCredentials struct {
	EmailOrUsername string `json:"emailOrUsername"`
	Password        string `json:"password"`
}

type Auth struct {
	Globals utils.Globals
}

func (auth Auth) SignIn() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var input LoginCredentials
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				api_errors.ErrorResponseBadPayloadFormat(
					err.Error(),
				),
			)
		}
		if len(input.EmailOrUsername) == 0 || len(input.Password) == 0 {
			return c.SendStatus(fiber.StatusUnauthorized)
		}

		checkInterpolation := "username = ?"
		if strings.ContainsRune(input.EmailOrUsername, '@') {
			checkInterpolation = "email = ?"
		}

		var user models.User

		if err := auth.Globals.DB.Model(&models.User{}).First(&user, checkInterpolation, input.EmailOrUsername).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendStatus(fiber.StatusUnauthorized)
			}
		}

		if user.Status == gormtypes.UserStatus_Pending {
			return c.Status(fiber.StatusUnauthorized).JSON(api_errors.Auth["unverified_email"])
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

		tokenString, expirationTime, err := jwt.GetJwt(auth.Globals, user)

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
				api_errors.ErrorResponseBadPayloadFormat(
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

		// Password encrypt
		password, _ := password.GenerateHashFromPassword(payload.Password, &password.Params{
			Memory:      64 * 1024,
			Iterations:  3,
			Parallelism: 2,
			SaltLength:  16,
			KeyLength:   32,
		})

		// Email confirmation token generation
		confirmationToken, _, err := jwt.GetConfirmationCodeToken(&auth.Globals, payload.Email)
		if err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		newUser := models.User{
			Id:               uuid.New().String(),
			Fname:            payload.Fname,
			Lname:            payload.Lname,
			Email:            payload.Email,
			Username:         payload.Username,
			BirthDate:        birthDate,
			Country:          payload.Country,
			Password:         password,
			ConfirmationCode: null.StringFrom(confirmationToken),
		}

		if err := auth.Globals.DB.Create(&newUser).Error; err != nil {
			// Read about error value: https://blog.golang.org/go1.13-errors#TOC_2.1.
			if pgErr, ok := err.(*pgconn.PgError); ok {
				switch pgErr.Code {
				// Unique constraint violation
				case "23505":
					var apiErrorCodeStatus api_errors.ApiErrorCodeStatus
					if pgErr.ConstraintName == "users_email_key" {
						apiErrorCodeStatus = api_errors.Auth["email_already_taken"]
					} else {
						apiErrorCodeStatus = api_errors.Auth["username_already_taken"]
					}
					return c.Status(fiber.StatusConflict).JSON(
						api_errors.ErrorResponseDataConflict(apiErrorCodeStatus),
					)
				}
			}
			auth.Globals.Logger.Fatal(
				"Error while creating a new user",
				zap.Error(err),
			)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		// TODO: Generate token for email confirmation

		// TODO: Add it into redis with an ETL to define
		// TODO: Send successful signup email with confirmation link

		id, err := auth.Globals.Emails.SendSignUpConfirmation(&emails.RecipientUser{
			Fname:            newUser.Fname,
			Lname:            newUser.Lname,
			Email:            newUser.Email,
			ConfirmationCode: confirmationToken,
		})

		if err != nil {
			auth.Globals.Logger.Fatal(
				"Error while sending sign-up confirmation email",
				zap.Error(err),
			)
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		auth.Globals.Logger.Info(
			"OK: sending sign-up confirmation email",
			zap.String("ID", id),
		)

		return c.Status(201).JSON(newUser.ToUserSafe())
	}
}

func (auth Auth) ConfirmEmail() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var user models.User

		if err := auth.Globals.DB.Where("status = ? AND confirmation_code = ?", string(gormtypes.UserStatus_Pending), c.Params("confirmationCode")).First(&user).Error; err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return c.SendStatus(fiber.StatusNotFound)
			}
		}

		updates := map[string]interface{}{
			"status":            string(gormtypes.UserStatus_Active),
			"confirmation_code": nil,
		}

		if err := auth.Globals.DB.Model(&user).Updates(updates).Error; err != nil {
			return c.SendStatus(fiber.StatusInternalServerError)
		}

		return c.SendStatus(fiber.StatusOK)
	}
}

/***
Restricted: need jwt
*/
func (auth Auth) LoggedUser() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		jwtUser, err := jwt.RetrieveUserFromCtx(c)

		if err != nil {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		return c.JSON(fiber.Map{"user": jwtUser})
	}
}
