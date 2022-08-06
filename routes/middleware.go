package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func NewFiber() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "[PID:${pid}] ${locals:requestid} [${ip}]:${port} ${status} - ${method} ${path} ${latency} ${error}\n",
	}))
	return app
}
