package discount

import (
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/tables/models"
)

func Discount(c *fiber.Ctx) error {
	fmt.Println(time.Now())
	db := database.DBConn
	quote := new(models.Quote)

	//parse the input json quote request
	if err := c.BodyParser(quote); err != nil {
		panic(err)
	}

	quote.RateTermStartDate = quote.Quotes.QuoteEffDate
	quote.RateAppliedDate = time.Now()
	quote.QuoteStartDate = time.Now()
	//Populate Good Coverages
	discresp := InvokeDiscRules(quote)
	db.Create(&quote)
	fmt.Println(time.Now())
	return c.JSON(discresp)
}

func InvokeDiscRules(quote *models.Quote) string {
	log.Println("unimplemented")
	panic("unimplemented")
}
