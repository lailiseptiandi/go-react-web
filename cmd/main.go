package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/template/html/v2"
	"github.com/lailiseptiandi/go-web-app/app/config"
	"github.com/lailiseptiandi/go-web-app/app/handlers"
	"github.com/lailiseptiandi/go-web-app/app/repository"
	"github.com/lailiseptiandi/go-web-app/app/services"
)

// var viewsfs embed.FS

func main() {
	config.LoadEnv()
	config.ConnectDB()

	engine := html.New("./web", ".html")
	engine.AddFunc("add", func(a, b int) int {
		return a + b
	})
	app := fiber.New(fiber.Config{
		Views: engine,
	},
	)

	app.Use(cors.New())

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	userRepo := repository.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	app.Get("/user", userHandler.GetUsers)
	app.Post("/user", userHandler.CreateUser)

	listen := fmt.Sprintf(":%s", os.Getenv("PORT"))
	app.Listen(listen)
}
