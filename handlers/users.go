package handlers

import (
	user "karma_files_go/helpers/users"
	"time"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type CreateUserRequest struct {
	Userid   string `json:"userid" form:"userid"`
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

	var date string = time.Now().String()

	rando, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm", 10)
	var api_token string = req.Userid + "---" + rando

	user.CreateUser(req.Userid, req.Password, date, api_token)
	return c.JSON(ResponseHTTP{
		Success: true,
		Message: "Successfully created user.",
		Data:    nil,
	})
}
