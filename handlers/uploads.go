package handlers

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gofiber/fiber/v2"

	awsutil "karma_files_go/aws"
	kf_config "karma_files_go/config"
	filespkg "karma_files_go/helpers/files"
)

type ResponseHTTP struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}

type CreateFileUserRequest struct {
	Name string     `json:"name" form:"name"`
	File *fiber.Ctx `json:"file" form:"file"`
}

func UploadSingleFile(c *fiber.Ctx) error {
	sdkConfig, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Couldn't load default configuration. Have you set up your AWS account?")
		fmt.Println(err)
	}
	s3Client := s3.NewFromConfig(sdkConfig)
	basics := awsutil.Bucket{
		S3Client: s3Client,
	}
	form, err := c.MultipartForm()
	// fid, _ := gonanoid.Generate("qwertyuiopasdfghjklzxcvbnm1234567890", 20)
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
	filename := parts[0]
	extension := parts[len(parts)-1]
	fid := filespkg.CreateFile(c.Locals("uid").(string), filename, "description")
	if fid == "" {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: "Error creating file",
		})
	}
	fileKey := fid + "." + extension
	if err := basics.UploadFile(kf_config.NewConfig().BuckerName, fileKey, fileKey); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
			Success: false,
			Data:    nil,
			Message: err.Error(),
		})
	}
	// if err := c.SaveFile(file, "./uploads/"+fid+"."+extension); err != nil {
	// 	return c.Status(fiber.StatusInternalServerError).JSON(ResponseHTTP{
	// 		Success: false,
	// 		Data:    nil,
	// 		Message: err.Error(),
	// 	})
	// }

	return c.JSON(ResponseHTTP{
		Success: true,
		Data:    nil,
		Message: "File uploaded successfully",
	})
}
