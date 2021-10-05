package routes

import (
	"app/controllers"
	"app/middlewares"
	"app/utils"
	"github.com/gofiber/fiber/v2"
)

// Setup Receiver to configure the Routes
func Setup(app *fiber.App) {

	// Group prefix /api/v1
	api := app.Group("api/v1")

	// Group prefix /api/v1/admin
	admin := api.Group("admin")

	// Login User - Public Route
	admin.Post("login", controllers.Login)

	// routes that needs to be authenticated.
	adminAuthenticated := admin.Use(middlewares.IsAuthenticated)

	// Group prefix /api/v1/admin/users
	adminUsers := adminAuthenticated.Group("users")

	// Get the authenticated user information
	adminAuthenticated.Post("user/logout", controllers.Logout)

	// ************** START User Paths **************
	// Create User /api/v1/admin/users/create
	adminUsers.Post("create", controllers.Create)
	// Get the authenticated user information /api/v1/admin/users/user
	adminUsers.Get("user/info", controllers.User)
	// Updated the user profile /api/v1/admin/users/user/update/info
	adminUsers.Put("user/update/info", controllers.UpdateInfo)
	// Updated the user password /api/v1/admin/users/user/update/password
	adminUsers.Put("user/update/password", controllers.UpdatePassword)
	// ************** END User Paths **************

	// Api Status
	api.Get("status", controllers.Status)

	// Handle 404
	app.Use(func(c *fiber.Ctx) error {
		err := c.SendStatus(fiber.StatusNotFound)
		utils.HandleError(err)
		return c.JSON(fiber.Map{
			"status":  fiber.StatusNotFound,
			"message": "Not Found",
		})
	})
}
