package controllers

import (
	"app/middlewares"
	"app/models"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"time"
)

// Login function to handle the login
func Login(c *fiber.Ctx) error {

	// parameters from the url will be store here.
	var data map[string]string

	// Getting the data from the params
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"data": data,
			"message": "error, data does not match!",
		})
	}

	// Creating the variable to store the user
	var user models.User

	// Getting the user from the database
	if !existsUser(data["email"], &user) {
		// Return HTTP 400
		c.Status(fiber.StatusBadRequest)
		// Return the message
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials!",
		})
	}

	// Check if password matches.
	if err := user.ComparePassword(data["password"]); err != nil {
		// Return HTTP 400
		c.Status(fiber.StatusBadRequest)
		// Return the message
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials!",
		})
	}

	// Variable to set the scope
	scope := "default"

	// Getting the token
	token, err := middlewares.GenerateJWT(user.Id, scope)

	// Check for errors
	if err != nil {
		fmt.Println("Token ", err)
		// Return HTTP 400
		c.Status(fiber.StatusBadRequest)
		// Return the message
		return c.JSON(fiber.Map{
			"message": "Invalid Credentials!",
		})
	}

	// Creating the Cookie
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}

	// Return the token as cookie
	c.Cookie(&cookie)

	// Just return the success message
	return c.JSON(fiber.Map{
		"message": "login successfully!",
	})
}

// Logout function to authenticate the user
func Logout(c *fiber.Ctx) error {

	// Setting the cookie with one hour ago the expires time.
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	// Setting the cookie
	c.Cookie(&cookie)

	// Just return the success message
	return c.JSON(fiber.Map{
		"message": "logout successfully",
	})
}
