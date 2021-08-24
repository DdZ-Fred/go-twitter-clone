package jwt

import (
	"os"
	"time"

	"github.com/DdZ-Fred/go-twitter-clone/models"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
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

// Results: tokenString, expirationTime, error
func GetJWT(globals utils.Globals, user models.User) (string, time.Time, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	expirationTime := time.Now().Add(time.Hour * 72)
	claims["authorized"] = true
	claims["user"] = JwtUser{
		Id:       user.Id,
		Fname:    user.Fname,
		Lname:    user.Lname,
		Email:    user.Email,
		Username: user.Username,
	}
	claims["aud"] = "go-twitter-clone"
	claims["iss"] = "golang-jwt"
	claims["exp"] = expirationTime.Unix()

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
