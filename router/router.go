package router

import (
	"github.com/buyanbadrakh/keycloak-group/config"
	"github.com/buyanbadrakh/keycloak-group/handler"
	"github.com/buyanbadrakh/keycloak-group/middleware"
	"github.com/buyanbadrakh/keycloak-group/utils"
	"github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	jwtware "github.com/gofiber/jwt/v3"
)

// var Handler *handler.Handler{}
// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// h := &handler.Handler{}
	SecretKey := "-----BEGIN CERTIFICATE-----\n" + config.Config.Keycloak.PublicKey + "\n-----END CERTIFICATE-----"
	key, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(SecretKey))

	app.Use(jwtware.New(jwtware.Config{
		SigningKey:     key,
		SigningMethod:  jwtware.RS256,
		ErrorHandler:   utils.JWTError,
		SuccessHandler: utils.JWTSuccess,
	}))

	app.Use(middleware.GetAccessToken)
	// Middleware
	// app.Use(middleware.Protected)
	api := app.Group(config.Config.Application.BasePath, logger.New())

	// Users
	users := api.Group("/users")
	users.Get("/", handler.GetUsers)
	users.Put("/:id/groups/:groupId", handler.GroupAddMember)
	users.Delete("/:id/groups/:groupId", handler.GroupDeleteMember)

	// Groups
	groups := api.Group("/groups")
	groups.Get("/", handler.GetGroups)
	groups.Get("/:id/members", handler.GetGroupMembers)
	groups.Post("/", handler.GroupCreate)
	groups.Put("/:id", handler.GroupUpdate)
	groups.Delete("/:id", handler.GroupDelete)
}
