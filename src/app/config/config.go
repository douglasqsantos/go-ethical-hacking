package config

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// DataBase struct
type DataBase struct {
	host     string
	port     string
	username string
	password string
	dbName   string
	dbOpts   string
	dbType   string
}

// Config struct
type Config struct {
	Timezone string
	DBType   string
	AppPort  string
	JwtKey   string
}

// bootStrap Receiver to bootstrap the App config
func (c *Config) bootStrap() {
	// Setting the Timezone
	timezone := os.Getenv("GO_APP_TZ")

	if timezone == "" {
		timezone = "Local"
	}

	// Setup the timezone
	c.Timezone = timezone

	// App Port
	appPort := os.Getenv("GO_APP_PORT")
	if appPort == "" {
		appPort = ":8000"
	} else {
		appPort = fmt.Sprintf(":%v", appPort)
	}

	// Setup the App port
	c.AppPort = appPort

	// JWT
	jwt := os.Getenv("GO_APP_JWT_KEY")
	if jwt == "" {
		c.JwtKey = "BM30qPhqBmRBUSwAavz0pvWfRDvCbXxNMl0JHfi6TpdkOqyEVjKisomnTnyvMqo"
	} else {
		c.JwtKey = jwt
	}

	// Getting the dbType
	dbType := os.Getenv("GO_DB_TYPE")

	// Getting the dbType
	c.DBType = strings.ToLower(dbType)
}

// New Function to return a pointer of a new Config
func New() *Config {
	// Creating an instance of Config
	cfg := &Config{}

	// Execute the bootstrap
	cfg.bootStrap()

	// Return a pointer to cfg
	return cfg
}

// SetConf Receiver to set all the DB Configuration
func (db *DataBase) setConf() {

	// Getting all the variables from the Environment
	db.host = os.Getenv("GO_DB_HOST")
	db.port = os.Getenv("GO_DB_PORT")
	db.username = os.Getenv("GO_DB_USERNAME")
	db.password = os.Getenv("GO_DB_PASSWORD")
	db.dbName = os.Getenv("GO_DB_NAME")
	db.dbOpts = os.Getenv("GO_DB_OPTS")
	db.dbType = os.Getenv("GO_DB_TYPE")

	// Validating the database type
	if (strings.ToLower(db.dbType) != "mysql") && (strings.ToLower(db.dbType) != "postgresql") {
		log.Fatalln("Database does not Have support Yet! Supported databases: [mysql, postgresql")
	}
}

// GetDSN Receiver to parse the DSN
// https://gorm.io/docs/connecting_to_the_database.html
func (db *DataBase) GetDSN() string {
	// Setting all the variables
	db.setConf()

	// Defining the DSN to store the DSN connection
	var dsn string

	// Return the DSN for MySQL
	if strings.ToLower(db.dbType) == "mysql" {
		// e.g: "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
		dsn = fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?%v", db.username, db.password, db.host, db.port, db.dbName, db.dbOpts)
	}

	// Return the DSN for PostgreSQL
	if strings.ToLower(db.dbType) == "postgresql" {
		// e.g:  "host=localhost user=gorm password=gorm port=5432 sslmode=disable TimeZone=America/Sao_Paulo"
		dsn = fmt.Sprintf("host=%v user=%v password=%v dbname=%v port=%v %v", db.host, db.username, db.password, db.dbName, db.port, db.dbOpts)
	}

	// Return DSN
	return dsn
}
