package controllers

import (
	"app/models"
	"github.com/gofiber/fiber/v2"
)

// Status receiver to return the application status
func Status(c *fiber.Ctx) error {
	// Use via model rather than maps
	status := models.Status{}
	status.Status = "Available"
	status.Environment = "development"
	status.Version = "0.0.1.rc"

	// Return the application status
	return c.JSON(status)
}
