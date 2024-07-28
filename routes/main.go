package routes

import (
	"github.com/gofiber/fiber/v2"

	handlers "karma_files_go/handlers"
	"karma_files_go/middlewares"
)

func SetupRoutes() *fiber.App {
	app := fiber.New()
	v1 := app.Group("/v1")

	// files
	files := v1.Group("/files")
	files.Post("/upload", middlewares.KFAPI, handlers.UploadSingleFile)
	files.Post("/uploadMultiple", handlers.UploadMultipleFiles)

	// users
	users := v1.Group("/users")
	users.Get("/", handlers.GetUsers)
	users.Post("/create", handlers.CreateUser)

	return app
}
