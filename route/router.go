package route

import (
	"go-fiber-gorm/handler"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(r *fiber.App) {
	r.Post("/api/v1/login", handler.LoginHandler)
	r.Get("/api/v1/users", handler.GetUserHandler)
	r.Post("/api/v1/users", handler.CreateUserHandler)
	r.Get("/api/v1/users/:id", handler.GetByIdUser)
	r.Put("/api/v1/users/:id", handler.UpdateByIdUser)
	r.Delete("/api/v1/users/:id", handler.DeleteByIdUser)

	r.Post("/api/v1/orders", handler.CreateOrderHandler)

}
