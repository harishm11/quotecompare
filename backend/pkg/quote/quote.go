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
	fmt.Println(time.Now())
	db := database.DBConn
	quote := new(models.Quote)

	//parse the input json quote request
	if err := c.BodyParser(quote); err != nil {
		panic(err)
	}

	quote.RateTermStartDate = quote.QuoteEffDate
	quote.RateAppliedDate = time.Now()
	quote.QuoteStartDate = time.Now()

	//Populate Coverages

	var c1 = models.Coverage{}
	var c2 = models.Coverage{}
	var c3 = models.Coverage{}
	var c4 = models.Coverage{}
	var c5 = models.Coverage{}
	var c6 = models.Coverage{}
	var c7 = models.Coverage{}

	c1 =
		models.Coverage{
			CoverageCode:       "BodilyInjury",
			CvgSymbol:          "",
			LimitPerPerson:     "100000",
			LimitPerOccurrence: "300000",
			Deductible:         "",
			CvgPremium:         0.0,
		}

	c2 =
		models.Coverage{
			CoverageCode:       "PropertyDamage",
			CvgSymbol:          "",
			LimitPerPerson:     "100000",
			LimitPerOccurrence: "",
			Deductible:         "",
			CvgPremium:         0.0,
		}
	c3 =
		models.Coverage{
			CoverageCode:       "Comprehensive",
			CvgSymbol:          "",
			LimitPerPerson:     "",
			LimitPerOccurrence: "",
			Deductible:         "500",
			CvgPremium:         0.0,
		}
	c4 =
		models.Coverage{
			CoverageCode:       "Medical",
			CvgSymbol:          "",
			LimitPerPerson:     "5000",
			LimitPerOccurrence: "",
			Deductible:         "",
			CvgPremium:         0.0,
		}
	c5 =
		models.Coverage{
			CoverageCode:       "Collision",
			CvgSymbol:          "",
			LimitPerPerson:     "",
			LimitPerOccurrence: "",
			Deductible:         "500",
			CvgPremium:         0.0,
		}
	c6 =
		models.Coverage{
			CoverageCode:       "UninsuredMotoristBI",
			CvgSymbol:          "",
			LimitPerPerson:     "100000",
			LimitPerOccurrence: "",
			Deductible:         "300000",
			CvgPremium:         0.0,
		}
	c7 =
		models.Coverage{
			CoverageCode:       "UninsuredMotoristPD",
			CvgSymbol:          "",
			LimitPerPerson:     "100000",
			LimitPerOccurrence: "",
			Deductible:         "",
			CvgPremium:         0.0,
		}
	for vehidx := range quote.Vehicles {
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c1)
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c2)
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c3)
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c4)
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c5)
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c6)
		quote.Vehicles[vehidx].Coverages = append(quote.Vehicles[vehidx].Coverages, c7)
	}

	//derive rating variables from the quote data
	plcyratvars := ratingvariables.PopPolicyRatingVars(quote)
	vehratvars := ratingvariables.PopVehicleRatingVars(quote.Vehicles)
	drvratvars := ratingvariables.PopDriverRatingVars(quote.Drivers)

	//call rating engine passing derived rating variables
	rtgengine.RatingEngineImpl(plcyratvars, drvratvars, vehratvars)

	db.Create(&quote)
	fmt.Println(time.Now())
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
