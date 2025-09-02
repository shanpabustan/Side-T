// package main

// import (
// 	"fmt"
// 	"log"
// 	//"intern_template_v1/controller"
// 	"intern_template_v1/middleware"
// 	"intern_template_v1/routes"

// 	"github.com/cloudinary/cloudinary-go/v2"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/gofiber/fiber/v2/middleware/cors"
// 	"github.com/gofiber/fiber/v2/middleware/logger"
// )

// func init() {
// 	fmt.Println("STARTING SERVER...")
// 	fmt.Println("INITIALIZE DB CONNECTION...")
// 	if middleware.ConnectDB() {
// 		fmt.Println("DB CONNECTION FAILED!")
// 	} else {
// 		fmt.Println("DB CONNECTION SUCCESSFUL!")
// 	}
// }

// // Function to initialize Cloudinary
// func initializeCloudinary() {
// 	// Load the Cloudinary URL from .env
// 	cloudinaryURL := middleware.GetEnv("CLOUDINARY_URL")
// 	if cloudinaryURL == "" {
// 		log.Fatal("Cloudinary URL is missing in the .env file.")
// 	}

// 	// Initialize Cloudinary with the URL from .env
// 	_, err := cloudinary.NewFromURL(cloudinaryURL)
// 	if err != nil {
// 		log.Fatalf("Error initializing Cloudinary: %v", err)
// 	}

// 	// Store Cloudinary client in global variable (optional, depends on your architecture)
// 	// If needed, you can store it in middleware or globally for access throughout your app.
// 	// For example:
// 	// middleware.CloudinaryClient = cld

// 	fmt.Println("Cloudinary initialized successfully!")
// }

// func main() {
// 	app := fiber.New(fiber.Config{
// 		AppName: middleware.GetEnv("PROJ_NAME"),
// 	})

// 	// CORS CONFIG - Move this before routes
// 	app.Use(cors.New(cors.Config{
// 		AllowOrigins: "*",
// 		AllowMethods: "GET,POST,PUT,DELETE",
// 		AllowHeaders: "Origin, Content-Type, Accept",
// 	}))
// 	// API ROUTES
// 	// Sample Endpoint
// 	// localhost:5566/check
// 	// app.Get("/check", controller.SampleController)

// 	// Do not remove this endpoint
// 	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
// 		return c.SendStatus(204) // No Content
// 	})

// 	routes.AppRoutes(app)

// 	// LOGGER
// 	app.Use(logger.New())

// 	// Start Server
// 	app.Listen(fmt.Sprintf(":%s", middleware.GetEnv("PROJ_PORT")))
// }


package main

import (
	"fmt"
	"log"

	"intern_template_v1/middleware"
	"intern_template_v1/routes"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	// Print startup messages
	fmt.Println("STARTING SERVER...")
	
	// Initialize DB connection
	fmt.Println("INITIALIZE DB CONNECTION...")
	if middleware.ConnectDB() {
		fmt.Println("DB CONNECTION FAILED!")
	} else {
		fmt.Println("DB CONNECTION SUCCESSFUL!")
	}

	// Initialize Cloudinary connection
	initializeCloudinary()
}

// Function to initialize Cloudinary
func initializeCloudinary() {
	// Load the Cloudinary URL from .env
	cloudinaryURL := middleware.GetEnv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		log.Fatal("Cloudinary URL is missing in the .env file.")
	}

	// Initialize Cloudinary with the URL from .env
	_, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		log.Fatalf("Error initializing Cloudinary: %v", err)
	}

	// Store Cloudinary client in global variable (optional, depends on your architecture)
	// If needed, you can store it in middleware or globally for access throughout your app.
	// For example:
	// middleware.CloudinaryClient = cld

	fmt.Println("Cloudinary initialized successfully!")
}

func main() {
	// Create Fiber app instance
	app := fiber.New(fiber.Config{
		AppName: middleware.GetEnv("PROJ_NAME"), // Set the project name from .env
	})

	// CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",                    // Allow all origins
		AllowMethods: "GET,POST,PUT,DELETE",  // Allow these HTTP methods
		AllowHeaders: "Origin, Content-Type, Accept", // Allow headers
	}))

	// Do not remove this endpoint
	// It serves a response to prevent 404 errors for favicon.ico requests
	app.Get("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendStatus(204) // No Content
	})

	// Register routes
	routes.AppRoutes(app)

	// Logging middleware
	app.Use(logger.New())

	// Start the server and listen on the configured port
	app.Listen(fmt.Sprintf(":%s", middleware.GetEnv("PROJ_PORT")))
}
