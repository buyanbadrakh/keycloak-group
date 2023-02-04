package handler

import (
	"encoding/json"
	"fmt"

	"github.com/buyanbadrakh/keycloak-group/config"
	"github.com/buyanbadrakh/keycloak-group/model"
	"github.com/buyanbadrakh/keycloak-group/utils"
	"github.com/gofiber/fiber/v2"
)

func GetUsers(c *fiber.Ctx) error {

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()
	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/users", config.Config.Keycloak.URL, user.Realm))
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	body, code, res := utils.CheckError(a)
	if body == nil {
		return c.Status(code).JSON(res)
	}

	var users []model.User
	if err := json.Unmarshal(body, &users); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "json unmarshal"})
	}

	var response []model.User
	for _, u := range users {
		for _, c := range u.Attributes.Company {
			if c == user.Company {
				response = append(response, u)
			}
		}
	}
	return c.JSON(response)
}

func GroupAddMember(c *fiber.Ctx) error {

	id := c.Params("id")
	groupId := c.Params("groupId")

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()
	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/users/%s/groups/%s", config.Config.Keycloak.URL, user.Realm, id, groupId))
	req.Header.SetMethod(fiber.MethodPut)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	body, code, res := utils.CheckError(a)
	if body == nil {
		return c.Status(code).JSON(res)
	}

	var response interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "json unmarshal"})
	}

	return c.JSON(fiber.Map{"code": code, "message": "added", "result": "added"})
}

func GroupDeleteMember(c *fiber.Ctx) error {

	id := c.Params("id")
	groupId := c.Params("groupId")

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()
	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/users/%s/groups/%s", config.Config.Keycloak.URL, user.Realm, id, groupId))
	req.Header.SetMethod(fiber.MethodDelete)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	body, code, res := utils.CheckError(a)
	if body == nil {
		return c.Status(code).JSON(res)
	}

	var response interface{}
	if err := json.Unmarshal(body, &response); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "json unmarshal"})
	}

	return c.JSON(fiber.Map{"code": code, "message": "deleted", "result": "deleted"})
}
