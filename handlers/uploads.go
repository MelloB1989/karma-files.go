package handlers

import (
	"github.com/gofiber/fiber/v2"
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
	if err := c.SaveFile(file, "./uploads/"+file.Filename); err != nil {
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