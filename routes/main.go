package routes

import (
	"github.com/gofiber/fiber/v2"
	handlers "karma_files_go/handlers"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/v1")

	//files
	files := v1.Group("/files")
	files.Post("/upload", handlers.UploadSingleFile)

	return app
}