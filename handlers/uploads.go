package handlers

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	gonanoid "github.com/matoous/go-nanoid/v2"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CreateUserRequest struct {
	Name	 string `json:"name" form:"name"`
	File	 *fiber.Ctx `json:"file" form:"file"`
}

func UploadSingleFile(c *fiber.Ctx) error {
	form, err := c.MultipartForm()
	fid, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890", 20);
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	files := form.File["file"]
	if len(files) == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "No file uploaded",
		})
	}

	file := files[0]
	parts := strings.Split(file.Filename, ".")
	extension := parts[len(parts)-1]
	if err := c.SaveFile(file, "./uploads/"+fid+"."+extension); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}

	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "File uploaded successfully",
	})
}