package quote

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
	"github.com/harishm11/quoteCompare/pkg/rtgengine"
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
	if err := c.BodyParser(quote); err != nil {
		panic(err)

	}

	quote.RateTermStartDate = quote.QuoteEffDate
	quote.RateAppliedDate = time.Now()
	quote.QuoteStartDate = time.Now()

	plcyratvars := ratingvariables.PopPolicyRatingVars(quote)
	vehratvars := ratingvariables.PopVehicleRatingVars(quote.Vehicles)
	drvratvars := ratingvariables.PopDriverRatingVars(quote.Drivers)

	rtgengine.RatingEngineImpl(quote)
	fmt.Println(plcyratvars)
	fmt.Println(vehratvars)
	fmt.Println(drvratvars)

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
