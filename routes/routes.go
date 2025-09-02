package routes

import (
	"intern_template_v1/controller"

	"github.com/gofiber/fiber/v2"
)

func AppRoutes(app *fiber.App) {
	// SAMPLE ENDPOINT
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Golang World!")
	})

	// CREATE YOUR ENDPOINTS HERE
	api := app.Group("/api")

	// api.Get("/users", controller.GetUsers)
	//api.Post("/users", controller.CreateUser)
	api.Post("/add",  controller.AddItem)
	api.Get("/fetch",controller.GetItems)
	api.Post("/buy",controller.AddPurchase)
	//api.Post("/approval", controller.CreateApprovalStatus)
	//api.Put("/users/:id/approve", controller.ApproveUser)
	// api.Put("/users/:id", controller.UpdateUser)
	// api.Delete("/users/:id", controller.DeleteUser)
	// --------------------------
}
