package utils

import "github.com/gofiber/fiber/v2"

func CheckError(a *fiber.Agent) ([]byte, int, fiber.Map) {
	if err := a.Parse(); err != nil {
		return nil, fiber.StatusInternalServerError, fiber.Map{"code": fiber.StatusInternalServerError, "message": "parse error"}
	}

	code, body, errs := a.InsecureSkipVerify().Bytes()

	if code != fiber.StatusOK || len(errs) > 0 {
		return nil, code, fiber.Map{"code": fiber.StatusInternalServerError, "message": "unknown_error"}
	}

	return body, 0, nil
}
