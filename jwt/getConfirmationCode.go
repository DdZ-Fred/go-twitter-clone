package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var JWT_EMAIL_CONFIRMATION_SIGNING_KEY = []byte(os.Getenv("JWT_EMAIL_CONFIRMATION_SIGNING_KEY"))

type ConfirmationCodeTokenCustomClaims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GetConfirmationCodeToken(globals *utils.Globals, userEmail string) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour * 24)

	claims := ConfirmationCodeTokenCustomClaims{
		userEmail,
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "golang-jwt",
			Audience:  "go-twitter-clone",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(JWT_EMAIL_CONFIRMATION_SIGNING_KEY)

	if err != nil {
		globals.Logger.Info("Error while retrieving signed token",
			zap.String("originalError", err.Error()),
			zap.Time("time", time.Now()),
		)
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

func VerifyConfirmationCodeToken(globals utils.Globals, confirmationCodeTokenStr string) (bool, string, error) {
	confirmationCodeToken, err := jwt.ParseWithClaims(confirmationCodeTokenStr, &ConfirmationCodeTokenCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return JWT_EMAIL_CONFIRMATION_SIGNING_KEY, nil
	})

	if err != nil {
		return false, "", err
	}

	if claims, ok := confirmationCodeToken.Claims.(*ConfirmationCodeTokenCustomClaims); ok && confirmationCodeToken.Valid {
		return true, claims.Email, nil
	}

	return false, "", fmt.Errorf("Invalid JWT Token")
}
