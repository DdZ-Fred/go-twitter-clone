package api

import (
	"github.com/DdZ-Fred/go-twitter-clone/handlers"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
)

func AuthRouter(globals utils.Globals) {
	authHandler := handlers.Auth{Globals: globals}
	authRouter := globals.App.Group("/auth")

	authRouter.Post("/login", authHandler.Login())
	authRouter.Post("/sign-up", authHandler.SignUp())
}
