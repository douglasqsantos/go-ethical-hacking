package controllers

import (
	"app/database"
	"app/middlewares"
	"app/models"
	"errors"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

// Create function to register the user
func Create(c *fiber.Ctx) error {

	// parameters from the url will be store here.
	var data map[string]string

	// Getting the data from the params
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// Creating the variable to store the user
	var user models.User

	// Check User
	if existsUser(data["email"], &user) {
		// Return HTTP 400
		c.Status(fiber.StatusBadRequest)
		// Return the message
		return c.JSON(fiber.Map{
			"message": "The email is already in use!",
		})
	}

	// Checking if the password matches
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	// Populating the User
	user.FirstName = data["first_name"]
	user.LastName = data["last_name"]
	user.Email = data["email"]

	// Creating a hash of the password
	user.SetPassword(data["password"])

	// Creating the User
	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error: %v\n", err.Error())
	}

	// Returning the user....
	return c.JSON(fiber.Map{
		"message": "User was created successfully",
	})
}

// User function to get authenticate user information
func User(c *fiber.Ctx) error {

	// Getting the user id
	id, _ := middlewares.GetUserId(c)

	// Setting the user variable
	var user models.User

	// Getting the user from the database
	database.DB.Where("id = ?", id).First(&user)

	// Return the user
	return c.JSON(user)

}

// UpdateInfo function to update the user profile
func UpdateInfo(c *fiber.Ctx) error {
	// parameters from the url will be store here.
	var data map[string]string

	// Getting the data from the params
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// Getting the user id
	id, _ := middlewares.GetUserId(c)

	// Setting the user information
	user := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}

	user.Id = id

	// Updating the user
	if err := database.DB.Model(&user).Updates(&user).Error; err != nil {
		log.Printf("Error: %v\n", err.Error())
	}

	// return the updated user
	return c.JSON(user)

}

// UpdatePassword function to update the user password
func UpdatePassword(c *fiber.Ctx) error {
	// parameters from the url will be store here.
	var data map[string]string

	// Getting the data from the params
	if err := c.BodyParser(&data); err != nil {
		return c.JSON(fiber.Map{
			"message": err,
		})
	}

	// Checking if the password matches
	if data["password"] != data["password_confirm"] {
		c.Status(400)
		return c.JSON(fiber.Map{
			"message": "passwords do not match",
		})
	}

	// Getting the user id
	id, _ := middlewares.GetUserId(c)

	// Setting the user information
	user := models.User{
		//Id: id,
	}
	// using in this way because the struct has embedded structs, and we cannot use id inside the struct.
	user.Id = id

	// Setting the user password
	user.SetPassword(data["password"])

	// Updating the user
	database.DB.Model(&user).Updates(&user)

	// Just return the success message
	return c.JSON(fiber.Map{
		"message": "Password was updated successfully",
	})

}

// Function to Check if the user exists or not
func existsUser(field string, user *models.User) bool {
	var err error
	// Checking if the email is already exist
	if err = database.DB.Where("email = ?", field).First(&user).Error; err != nil {
		log.Printf("Warn: Checking %v\n", err.Error())
	}

	// Check if the requested user is already registered
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return true
	} else {
		return false
	}
}
