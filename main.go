package main

import (
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/pandeyvijay99/gohummbird-teacher/db"
	"github.com/pandeyvijay99/gohummbird-teacher/router"
)

func main() {
	app := fiber.New()
	app.Use(cors.New())

	db.ConnectDB()

	router.SetupRoutes(app)
	app.Listen(3000)

	defer db.DB.Close()
}
