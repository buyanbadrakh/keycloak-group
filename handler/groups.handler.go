package handler

import (
	"encoding/json"
	"fmt"

	"github.com/buyanbadrakh/keycloak-group/config"
	"github.com/buyanbadrakh/keycloak-group/model"
	"github.com/buyanbadrakh/keycloak-group/utils"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func GetGroups(c *fiber.Ctx) error {

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()

	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/groups?briefRepresentation=false", config.Config.Keycloak.URL, user.Realm))
	req.Header.SetMethod(fiber.MethodGet)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	body, code, res := utils.CheckError(a)
	if body == nil {
		return c.Status(code).JSON(res)
	}

	var groups []model.Group
	if err := json.Unmarshal(body, &groups); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"code": fiber.StatusInternalServerError, "message": "json unmarshal"})
	}

	var response []model.Group
	for _, g := range groups {
		for _, c := range g.Attributes.Company {
			if c == user.Company {
				response = append(response, g)
			}
		}
	}
	return c.JSON(fiber.Map{"code": "SUCCESS", "info": nil, "key": 0, "result": response})
}
func GetGroupMembers(c *fiber.Ctx) error {

	id := c.Params("id")

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()

	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/groups/%s/members", config.Config.Keycloak.URL, user.Realm, id))
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

	// return c.JSON(users)
	return c.JSON(fiber.Map{"code": "SUCCESS", "info": nil, "key": 0, "result": users})
}

func GroupCreate(c *fiber.Ctx) error {

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()

	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/groups", config.Config.Keycloak.URL, user.Realm))
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(fiber.MethodPost)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	var group model.Group
	json.Unmarshal(c.Body(), &group)

	group.Attributes = model.GroupAttr{
		Company: []string{user.Company},
	}

	b, _ := json.Marshal(group)

	log.Info("GroupCreate [REQUEST]: ", string(b))
	req.AppendBody(b)

	_, code, res := utils.CheckError(a)
	if code >= fiber.StatusBadRequest {
		return c.Status(code).JSON(res)
	}

	return c.JSON(fiber.Map{"code": "SUCCESS", "message": "created"})
}

func GroupUpdate(c *fiber.Ctx) error {

	id := c.Params("id")

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()

	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/groups/%s", config.Config.Keycloak.URL, user.Realm, id))
	// req.UseHostHeader = false
	req.Header.SetContentType("application/json")
	req.Header.SetMethod(fiber.MethodPut)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	var group model.Group
	json.Unmarshal(c.Body(), &group)

	group.ID = &id
	group.Attributes = model.GroupAttr{
		Company: []string{user.Company},
	}

	b, _ := json.Marshal(group)
	log.Info("GroupUpdate [REQUEST]: ", string(b))
	req.AppendBody(b)
	_, code, res := utils.CheckError(a)
	if code >= fiber.StatusBadRequest {
		return c.Status(code).JSON(res)
	}

	return c.JSON(fiber.Map{"code": "SUCCESS", "message": "updated"})
}

func GroupDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	user := c.Locals("user").(model.AccessModel)
	token := c.Locals("token").(model.Token)

	a := fiber.AcquireAgent()

	req := a.Request()
	req.SetRequestURI(fmt.Sprintf("%s/admin/realms/%s/groups/%s", config.Config.Keycloak.URL, user.Realm, id))
	req.Header.SetMethod(fiber.MethodDelete)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token.AccessToken))

	_, code, res := utils.CheckError(a)
	if code >= fiber.StatusBadRequest {
		return c.Status(code).JSON(res)
	}

	return c.JSON(fiber.Map{"code": "SUCCESS", "message": "deleted"})
}
