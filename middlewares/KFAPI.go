package middlewares

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"

	"karma_files_go/utils"
)

func KPAPI(c *fiber.Ctx) error {
	authHeader := c.GetReqHeaders()
	// Check if the header exists and has at least one value
	if values, exists := authHeader["Authorization"]; exists && len(values) > 0 {
		// Split the value by spaces and select the token part
		parts := strings.Split(values[0], " ")
		if len(parts) == 2 && parts[0] == "Bearer" {
			token := parts[1]
			decoded, err := utils.Decode(token)
			if err != nil {
				fmt.Println("Error decoding token:", err)
			}
			c.Locals("email", decoded["email"])
			c.Locals("uid", decoded["uid"])
			return c.Next()
		} else {
			fmt.Println("Invalid Authorization header format")
			return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
		}
	}
	// If the header does not exist, is empty, or does not match "Bearer <token>", handle accordingly
	return c.Status(fiber.StatusUnauthorized).SendString("Unauthorized")
}
