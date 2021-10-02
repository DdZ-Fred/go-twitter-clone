package api

import (
	"github.com/DdZ-Fred/go-twitter-clone/handlers"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
)

func AuthRouter(globals utils.Globals) {
	authHandler := handlers.Auth{Globals: globals}
	authRouter := globals.App.Group("/auth")

	authRouter.Post("/sign-in", authHandler.SignIn())
	authRouter.Post("/sign-up", authHandler.SignUp())
	authRouter.Get("/google-oauth-code-challenge", authHandler.GoogleCodeChallenge())
	authRouter.Post("/google-oauth-token", authHandler.GoogleOauthToken())
}
