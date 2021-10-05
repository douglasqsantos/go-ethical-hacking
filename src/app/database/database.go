package database

import (
	"app/config"
	"app/models"
	"app/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"strings"
)

// DB Using as global variable
var DB *gorm.DB

// dbSQLConnect function to database connect
func dbSQLConnect(dbType string) {
	// Errors
	var err error

	// Getting an instance of dbConfig
	dbConfig := config.DataBase{}

	// Getting the DSN
	dsn := dbConfig.GetDSN()

	// Defining the default message
	msg := fmt.Sprintf("Could not connect with the database [%v]", strings.ToTitle(dbType))

	switch {
	case dbType == "mysql":
		// Open the connection
		DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		utils.PanicMSG(err, msg)
	case dbType == "postgresql":
		// Open the connection
		DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		utils.PanicMSG(err, msg)
	default:
		log.Fatalln("Database does not have support Yet! Supported Databases: [MySQL and PostgreSQL")
	}
}

// Connect function to connect to the database and set the DB
func Connect() {
	// Getting the Database type
	cfg := config.New()

	fmt.Println("CFG --> ", cfg)

	// Setting up the Database
	switch {
	case cfg.DBType == "mysql":
		log.Println("Using MySQL Database")
		dbSQLConnect(cfg.DBType)
	case cfg.DBType == "postgresql":
		log.Println("Using PostgreSQL Database")
		dbSQLConnect(cfg.DBType)
	default:
		log.Fatalln("Database does not have support Yet! Supported Databases: [MySQL and PostgreSQL]")
	}
}

// AutoMigrate function to create the migrations
func AutoMigrate() {
	// Migrate the models
	err := DB.AutoMigrate(models.User{})
	// Check errors
	utils.HandleError(err)
}
