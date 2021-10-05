package middlewares

import (
	"app/config"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Variables to generate the JWT
var (
	cfg        *config.Config = config.New()
	SecretKey                 = cfg.JwtKey
	CookieName                = "jwt"
)

// ClaimsWithScope struct to control the claims for the JWT
type ClaimsWithScope struct {
	jwt.StandardClaims
	Scope string
}

// IsAuthenticated function to check if the user is authenticated or not
func IsAuthenticated(c *fiber.Ctx) error {
	// getting the cookie
	cookie := c.Cookies(CookieName)

	// getting the token
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// check if we have errors and if the token is not valid.
	if err != nil || !token.Valid {
		// Return HTTP 401
		c.Status(fiber.StatusUnauthorized)
		// Return the message
		return c.JSON(fiber.Map{
			"message": "Unauthenticated!",
		})
	}

	// if the user is authenticated it will send to the next request.
	return c.Next()
}

// GenerateJWT function to generate the JWT
func GenerateJWT(id uint, scope string) (string, error) {
	// payload to populate with the JWT
	payload := ClaimsWithScope{}
	payload.Subject = strconv.Itoa(int(id))
	payload.ExpiresAt = time.Now().Add(time.Hour * 24).Unix()
	payload.Scope = scope

	// Generating the token
	return jwt.NewWithClaims(jwt.SigningMethodHS256, payload).SignedString([]byte(SecretKey))
}

// GetUserId function to get the user id from the cookie
func GetUserId(c *fiber.Ctx) (uint, error) {
	// getting the cookie
	cookie := c.Cookies(CookieName)

	// getting the token
	token, err := jwt.ParseWithClaims(cookie, &ClaimsWithScope{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	// check if we have errors if it does return 0 and the error
	if err != nil {
		return 0, err
	}

	// Getting the user payload
	payload := token.Claims.(*ClaimsWithScope)

	// convert the string into uint
	id, _ := strconv.Atoi(payload.Subject)

	// return the id and the error.
	return uint(id), nil
}
