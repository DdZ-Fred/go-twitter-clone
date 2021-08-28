package apierrors

import "github.com/gofiber/fiber/v2"

type ApiErrorCodeStatus struct {
	Code   string
	Status string
}

func ErrorResponseBadPayloadFormat(message string) fiber.Map {
	return fiber.Map{
		"type":    "bad-payload-format",
		"message": message,
	}
}

func ErrorResponseDataConflict(apiErrorCodeStatus ApiErrorCodeStatus) fiber.Map {
	return fiber.Map{
		"type":   "data-conflict",
		"code":   apiErrorCodeStatus.Code,
		"status": apiErrorCodeStatus.Status,
	}
}
