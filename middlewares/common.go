package middlewares

import (
	"github.com/DdZ-Fred/go-twitter-clone/jwt"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
)



func InitMiddlewares() *utils.Middlewares {
	return &utils.Middlewares{
		JwtAuth: jwt.JwtMiddleware(),
	}
}