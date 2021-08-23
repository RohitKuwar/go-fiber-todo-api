package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/RohitKuwar/go-fiber-api/controllers"
)

func TodoRoute(route fiber.Router) {
	route.Get("/", controllers.GetTodos)
	route.Post("/", controllers.CreateTodo)
	route.Put("/:id", controllers.UpdateTodo)
	route.Delete("/:id", controllers.DeleteTodo)
	route.Get("/:id", controllers.GetTodo)
}

