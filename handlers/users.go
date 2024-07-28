package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"

	user "karma_files_go/helpers/users"
)

type CreateUserRequest struct {
	Userid   string `json:"userid"   form:"userid"`
	Password string `json:"password" form:"password"`
}

func GetUsers(c *fiber.Ctx) error {
	users, err := user.GetUsers()
	if err == nil {
		return c.JSON(ResponseHTTP{
			Success: true,
			Message: "Successfully retrieved all users.",
			Data:    users,
		})
	} else {
		return c.JSON(ResponseHTTP{
			Success: true,
			Message: "Failed to retrieve all users.",
			Data:    nil,
		})
	}
}

func CreateUser(c *fiber.Ctx) error {
	req := new(CreateUserRequest)
	if err := c.BodyParser(req); err != nil {
		return c.Status(400).JSON(ResponseHTTP{
			Success: false,
			Message: "Failed to parse request body.",
			Data:    nil,
		})
	}

	date := time.Now().Format("2006-01-02")
	// date := "2024-07-27"

	rando, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm", 10)
	api_token := req.Userid + "---" + rando

	user.CreateUser(req.Userid, req.Password, date, api_token)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}
