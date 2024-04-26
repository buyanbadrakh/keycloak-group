package utils

import (
	"encoding/json"

	"github.com/gofiber/fiber/v2/log"

	"github.com/gofiber/fiber/v2"
)

func CheckError(a *fiber.Agent) ([]byte, int, fiber.Map) {
	if err := a.Parse(); err != nil {
		return nil, fiber.StatusInternalServerError, fiber.Map{"code": fiber.StatusInternalServerError, "message": "parse error"}
	}

	code, body, errs := a.InsecureSkipVerify().Bytes()

	log.Info(string(body))
	if code >= fiber.StatusBadRequest || len(errs) > 0 {
		if body != nil {
			var b fiber.Map
			json.Unmarshal(body, &b)
			return nil, code, fiber.Map{"code": code, "message": b["error"]}
		}
		return nil, code, fiber.Map{"code": fiber.StatusInternalServerError, "message": "unknown_error"}
	}

	return body, code, nil
}
