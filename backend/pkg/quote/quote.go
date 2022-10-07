package quote

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/tables/models"
)

func GetQuotes(c *fiber.Ctx) error {
	db := database.DBConn
	var quotes []models.Quote
	result := db.Find(&quotes)
	if result.Error != nil {
		panic(result.Error)
	}
	return c.JSON(quotes)
}

func GetQuote(c *fiber.Ctx) error {
	id := c.Params("QuoteNumber")
	db := database.DBConn
	var quote models.Quote
	result := db.Find(&quote, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return c.JSON(quote)
}

func NewQuote(c *fiber.Ctx) error {

	db := database.DBConn
	quote := new(models.Quote)

	//parse the input json quote request
	if err := c.BodyParser(quote); err != nil {
		panic(err)
	}
	quote.RateTermStartDate = quote.QuoteEffDate
	quote.RateAppliedDate = time.Now()
	quote.QuoteStartDate = time.Now()

	db.Create(&quote)
	return c.JSON(quote)
}

func DeleteQuote(c *fiber.Ctx) error {
	id := c.Params("QuoteNumber")
	db := database.DBConn

	var quote models.Quote
	result := db.First(&quote, id)
	if result.Error != nil {
		panic(result.Error)
	}

	result1 := db.Delete(&quote)
	if result1.Error != nil {
		panic(result1.Error)
	}
	return c.SendString("Deleted")
}
