package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/harishm11/quoteCompare/database"
	quote "github.com/harishm11/quoteCompare/quote"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

func Welcome(c *fiber.Ctx) {
	c.Send("Hello !")
}

func setupRoutes(app *fiber.App) {
	app.Get("/", Welcome)
	app.Get("/quoteApi/quote", quote.GetQuotes)
	app.Get("/quoteApi/quote/:quoteNumber", quote.GetQuote)
	app.Post("/quoteApi/quote", quote.NewQuote)
	app.Delete("/quoteApi/quote/:quoteNumber", quote.DeleteQuote)
}

func initDatabase() {
	var err error
	database.DBConn, err = gorm.Open("sqlite3", "quote.db")
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")
	database.DBConn.AutoMigrate(&quote.Quote{})
	fmt.Println("Database Migrated")

}

func main() {
	app := fiber.New()
	initDatabase()
	defer database.DBConn.Close()
	setupRoutes(app)
	app.Listen(3000)

}
