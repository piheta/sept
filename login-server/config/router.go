package config

import (
	"errors"
	"os"

	json "github.com/goccy/go-json"
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/swagger"
	"github.com/piheta/sept/login-server/controllers"
	_ "github.com/piheta/sept/login-server/docs"
	weberrors "github.com/piheta/sept/login-server/errors"
)

func NewRouter(
	authController *controllers.AuthController,
	userController *controllers.UserController,
) *fiber.App {
	app := fiber.New(fiber.Config{
		JSONEncoder: json.Marshal,
		JSONDecoder: json.Unmarshal,

		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			// Retrieve the custom status code if it's a *WebError
			var e2 *weberrors.WebError
			if errors.As(err, &e2) {
				code = e2.StatusCode
			}
			c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
	})
	app.Use(logger.New())
	app.Use(recover.New())
	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")
	// Public Routes
	api.Post("/login", authController.Login)
	api.Post("/users/", userController.CreateUser)

	// JWT Middleware
	api.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return fiber.NewError(fiber.StatusUnauthorized, "invalid or expired jwt")
		},
	}))

	// Restricted Routes
	users := api.Group("/users")
	users.Get("", userController.GetUsers)
	// users.Post("", userController.CreateUser)
	users.Get("/:id", userController.GetUser)
	users.Put("/:id", userController.UpdateUser)
	users.Delete("/:id", userController.DeleteUser)

	return app
}
