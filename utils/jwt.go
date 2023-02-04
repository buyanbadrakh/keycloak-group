package utils

import (
	"encoding/json"
	"strings"

	"github.com/buyanbadrakh/keycloak-group/model"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

func JWTError(c *fiber.Ctx, err error) error {
	if err.Error() == "Missing or malformed JWT" {
		return c.Status(fiber.StatusBadRequest).
			JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT"})
	}

	return c.Status(fiber.StatusUnauthorized).
		JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
}

func JWTSuccess(c *fiber.Ctx) error {

	t := c.Locals("user").(*jwt.Token)
	// fmt.Println(user)
	claims, ok := t.Claims.(jwt.MapClaims)
	if ok && t.Valid {
		var a model.AccessModel
		jsonbody, err := json.Marshal(claims)
		if err != nil {
			return err
		}

		if err := json.Unmarshal(jsonbody, &a); err != nil {
			return err
		}

		a.Realm = a.Iss[strings.LastIndex(a.Iss, "/")+1:]

		c.Locals("user", a)
		c.Next()
		return nil
	} else {
		return c.Status(fiber.StatusUnauthorized).
			JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
	}

	// return c.Status(fiber.StatusUnauthorized).
	// 	JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWT"})
}
