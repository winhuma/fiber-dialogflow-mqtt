package routes

import (
	"fiber-dialogflow-mqtt/handlers"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(app *fiber.App) {
	gRoutes := app.Group("/")
	routeOther(gRoutes)
	routeLineBot(gRoutes)
}

func routeOther(g fiber.Router) {
	g.Get("", handlers.Hello)
}

func routeLineBot(g fiber.Router) {
	g.Post("callback", handlers.LineBotCallback)
}
