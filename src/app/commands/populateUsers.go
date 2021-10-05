package main

import (
	"app/database"
	"app/models"
	"log"
)

// Running the file
// docker compose exec api sh
// go run commands/populateUsers.go
// docker exec -it app_api_1 go run commands/populateUsers.go

func main() {
	// Connect to Database
	database.Connect()

	// Creating a default admin
	user := models.User{
		FirstName: "Douglas",
		LastName:  "Quintiliano dos Santos",
		Email:     "admin@admin.com",
		IsAdmin:   true,
	}

	// Setting up the password
	user.SetPassword("password")

	// Creating the user
	if err := database.DB.Create(&user).Error; err != nil {
		log.Printf("Error: %v\n", err.Error())
	}

}
