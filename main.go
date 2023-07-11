package main

import (
	"clean-architecture/repository"
	"clean-architecture/service"
	"clean-architecture/exception"
	"clean-architecture/controller"
	"clean-architecture/config"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main()  {
	//setup configuration
	configuration := config.New()
	database := config.NewMongoDatabase(configuration)

	//setup repository
	productRepository := repository.NewProductRepository(database)

	//setup service
	productService := service.NewProductService(&productRepository)

	// Setup Controller
	productController := controller.NewProductController(&productService)

	// Setup Fiber
	app := fiber.New(config.NewFiberConfig())
	app.Use(recover.New())

	// Setup Routing
	productController.Route(app)

	// Start App
	err := app.Listen(":3000")
	exception.PanicIfNeeded(err)

}