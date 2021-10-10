package handlers

import (
	"fmt"
	"os"

	"github.com/DdZ-Fred/go-twitter-clone/api_errors"
	"github.com/DdZ-Fred/go-twitter-clone/utils"
	"github.com/gofiber/fiber/v2"
)

/**
  GoogleOauthTokenPayload:
  - 'form' annotations are for application/x-www-form-urlencoded data mappings.
    Needed by c.BodyParser to properly parse the data
**/
type GoogleOauthTokenPayload struct {
	Code         string `json:"code" form:"code"`
	ClientId     string `json:"client_id" form:"client_id"`
	RedirectUri  string `json:"redirect_uri" form:"redirect_uri"`
	ResponseType string `json:"response_type" form:"response_type"`
	Audience     string `json:"audience" form:"audience"`
	GrantType    string `json:"grant_type" form:"grant_type"`
	CodeVerifier string `json:"code_verifier" form:"code_verifier"`
}

type GoogleOauthTokenEndpointResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	IdToken      string `json:"id_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	RefreshToken string `json:"refresh_token"`
}

func (auth Auth) GoogleCodeChallenge() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		codeVerifier := utils.GenerateCodeVerifier()
		return c.JSON(fiber.Map{
			"codeVerifier":  codeVerifier.Value,
			"codeChallenge": codeVerifier.CodeChallengeS256(),
		})
	}
}

/**
  GoogleOauthToken:

  - Content-Type: application/x-www-form-urlencoded
**/
func (auth Auth) GoogleOauthToken() func(*fiber.Ctx) error {
	return func(c *fiber.Ctx) error {
		var payload GoogleOauthTokenPayload

		if err := c.BodyParser(&payload); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				api_errors.ErrorResponseBadPayloadFormat(
					err.Error(),
				),
			)
		}

		body := fmt.Sprintf("code=%s&client_id=%s&client_secret=%s&redirect_uri=%s&grant_type=%s&code_verifier=%s",
			payload.Code,
			payload.ClientId,
			os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET"),
			payload.RedirectUri,
			payload.GrantType,
			payload.CodeVerifier,
		)

		resp, err := auth.Globals.RestyClient.R().
			SetHeader("Content-Type", "application/x-www-form-urlencoded").
			SetBody(body).
			SetResult(&GoogleOauthTokenEndpointResponse{}).
			Post("https://oauth2.googleapis.com/token")

		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(
				api_errors.ErrorResponseBadPayloadFormat(
					err.Error(),
				),
			)
		}

		if resp.IsError() {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"statusCode": resp.StatusCode(),
				"body":       resp.Result(),
			})
		}

		result := resp.Result().(*GoogleOauthTokenEndpointResponse)

		return c.Status(fiber.StatusOK).JSON(result)
	}
}
