package main

import (
	"os"
	"twb/controller"
	"twb/lib"
	"twb/repository"
	"twb/service"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	app := fiber.New()
	app.Use(cors.New(cors.Config{
		AllowOrigins: os.Getenv("FRONTEND_URL"),
		AllowMethods: "GET, POST",
	}))

	lib.InitRds()

	repository := repository.NewUploadRepo(lib.Rds)
	service := service.NewUploadService(repository)
	controller := controller.NewUploadController(service)

	app.Static("/uploads", "./uploads")
	upload := app.Group("api/v1/upload")

	upload.Post("/", controller.Post)
	upload.Get("/:id", controller.GetById)

	app.Listen(":" + os.Getenv("APP_PORT"))
}
