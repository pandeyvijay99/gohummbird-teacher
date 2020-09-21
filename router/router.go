package router

import (
	"github.com/gofiber/fiber"
	"github.com/gofiber/logger"
	"github.com/pandeyvijay99/gohummbird-teacher/handlers"
	"github.com/pandeyvijay99/gohummbird-teacher/middleware"
)

// SetupRoutes setup router api
func SetupRoutes(app *fiber.App) {
	// Middleware
	api := app.Group("/api", logger.New())
	api.Get("/", handlers.Hello)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)

	// User
	//user := api.Group("/user", middleware.Protected())
	//user.Get("/", handlers.GetUsers)
	//user.Get("/:id", handlers.GetUser)
	//user.Post("/", handlers.CreateUser)
	//user.Patch("/:id", handlers.UpdateUser)
	//user.Delete("/:id", handlers.DeleteUser)

	// Product
	product := api.Group("/product", middleware.Protected())
	product.Get("/", handlers.GetAllProducts)
	product.Get("/:id", handlers.GetProduct)
	product.Post("/", handlers.CreateProduct)
	product.Delete("/:id", handlers.DeleteProduct)



// Teacher
	teacher := api.Group("/teacher")
	teacher.Post("/create", handlers.AddNewTeacher)
//	teacher.Post("/update/:id", handlers.UpdateTeacher)
	teacher.Post("/create/:emp_id/:type", handlers.UpdateTeacher)
	teacher.Get("/get-all", handlers.GetAllTeachers)
	teacher.Get("/get/:id", handlers.GetTeacher)
	teacher.Patch("/details/:id",handlers.StoreBankDetails)


}
