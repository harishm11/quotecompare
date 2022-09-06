package urls

import (
	"github.com/gofiber/fiber/v2"
	quote "github.com/harishm11/quoteCompare/pkg/views"
)

func Home(c *fiber.Ctx) error {
	return c.SendString("Home")
}

func SetupRoutes(app *fiber.App) {
	app.Get("/", Home)
	app.Get("/quoteApi/quote", quote.GetQuotes)
	app.Get("/quoteApi/quote/:quoteNumber", quote.GetQuote)
	app.Post("/quoteApi/quote", quote.NewQuote)
	app.Delete("/quoteApi/quote/:quoteNumber", quote.DeleteQuote)
}
