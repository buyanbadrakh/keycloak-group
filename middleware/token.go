package middleware

import (
	"encoding/json"
	"fmt"

	"github.com/buyanbadrakh/keycloak-group/config"
	"github.com/buyanbadrakh/keycloak-group/model"
	"github.com/gofiber/fiber/v2"
)

func GetAccessToken(c *fiber.Ctx) error {

	a := fiber.AcquireAgent()

	// a.TLSConfig(config.TLSConfig)

	args := fiber.AcquireArgs()
	args.Set("grant_type", config.Config.Keycloak.GrantType)
	args.Set("client_id", config.Config.Keycloak.ClientID)
	args.Set("client_secret", config.Config.Keycloak.ClientSecret)

	a.Form(args)

	defer fiber.ReleaseAgent(a)
	defer a.ConnectionClose()

	req := a.Request()
	req.Header.SetMethod("POST")
	req.SetRequestURI(fmt.Sprintf("%s/realms/master/protocol/openid-connect/token", config.Config.Keycloak.URL))
	req.Header.Add("Connection", "close")

	if err := a.Parse(); err != nil {
		return c.Status(fiber.StatusInternalServerError).
			JSON(fiber.Map{"status": "error", "message": err.Error()})
	}
	code, body, errs := a.InsecureSkipVerify().Bytes()

	if code != fiber.StatusOK && errs != nil {
		return c.Status(code).
			JSON(fiber.Map{"status": "error", "message": errs[0].Error()})
	} else {
		token := model.Token{}

		if err := json.Unmarshal(body, &token); err != nil {
			return c.Status(code).
				JSON(fiber.Map{"status": "error", "message": err.Error()})
		}
		c.Locals("token", token)
		return c.Next()
	}
}
