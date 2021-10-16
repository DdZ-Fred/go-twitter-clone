package jwt

import (
	"fmt"
	"os"
	"time"

	"github.com/DdZ-Fred/go-twitter-clone/models"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

var JWT_SIGNING_KEY = []byte(os.Getenv("JWT_SIGNING_KEY"))

type JwtUser struct {
	Id       string `json:"id"`
	Fname    string `json:"fname"`
	Lname    string `json:"lname"`
	Email    string `json:"email"`
	Username string `json:"username"`
}

type CustomClaims struct {
	User JwtUser `json:"user"`
	jwt.StandardClaims
}

// Results: tokenString, expirationTime, error
func GetJwt(globals utils.Globals, user models.User) (string, time.Time, error) {
	expirationTime := time.Now().Add(time.Hour * 72)

	claims := CustomClaims{
		JwtUser{
			Id:       user.Id,
			Fname:    user.Fname,
			Lname:    user.Lname,
			Email:    user.Email,
			Username: user.Username,
		},
		jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			Issuer:    "golang-jwt",
			Audience:  "go-twitter-clone",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(JWT_SIGNING_KEY)
	if err != nil {
		globals.Logger.Info("Error while retrieving signed token",
			zap.String("originalError", err.Error()),
			zap.Time("time", time.Now()),
		)
		return "", time.Now(), err
	}

	return tokenString, expirationTime, nil
}

// func RetrieveJwtUser(globals utils.Globals, tokenString string) (*JwtUser, error) {
// 	userToken, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return JWT_SIGNING_KEY, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	if claims, ok := userToken.Claims.(*CustomClaims); ok && userToken.Valid {
// 		return &claims.User, nil
// 	}

// 	return nil, fmt.Errorf("Invalid JWT Token")
// }

// Depends on fiber jwt middleware
func RetrieveUserFromCtx(c *fiber.Ctx) (*JwtUser, error) {
	userToken := c.Locals("user").(*jwt.Token)
	if claims, ok := userToken.Claims.(*CustomClaims); ok {
		return &claims.User, nil
	}
	return nil, fmt.Errorf("Invalid JWT Token")
}
