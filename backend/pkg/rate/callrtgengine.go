package rate

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/harishm11/quoteCompare/contracts"
	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
	"github.com/harishm11/quoteCompare/pkg/rtgengine"
	"github.com/harishm11/quoteCompare/tables/models"
)

func Rate(c *fiber.Ctx) error {
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
	//Populate Good Coverages
	rateresp := InvokeRtgEngine(quote)
	db.Create(&quote)
	fmt.Println(time.Now())
	return c.JSON(rateresp)
}

func InvokeRtgEngine(quote *models.Quote) *contracts.RateResponse {
	var rateresp = new(contracts.RateResponse)
	var c1 = models.Coverage{CoverageCode: "BodilyInjury", CvgSymbol: "1", LimitPerPerson: "100000", LimitPerOccurrence: "300000", Deductible: "", CvgPremium: 0.0}
	var c2 = models.Coverage{CoverageCode: "PropertyDamage", CvgSymbol: "1", LimitPerPerson: "100000", LimitPerOccurrence: "", Deductible: "", CvgPremium: 0.0}
	var c3 = models.Coverage{CoverageCode: "Comprehensive", CvgSymbol: "1", LimitPerPerson: "", LimitPerOccurrence: "", Deductible: "500", CvgPremium: 0.0}
	var c4 = models.Coverage{CoverageCode: "Medical", CvgSymbol: "1", LimitPerPerson: "5000", LimitPerOccurrence: "", Deductible: "", CvgPremium: 0.0}
	var c5 = models.Coverage{CoverageCode: "Collission", CvgSymbol: "1", LimitPerPerson: "", LimitPerOccurrence: "", Deductible: "500", CvgPremium: 0.0}
	var c6 = models.Coverage{CoverageCode: "UninsuredMotoristBI", CvgSymbol: "1", LimitPerPerson: "100000", LimitPerOccurrence: "300000", Deductible: "", CvgPremium: 0.0}
	var c7 = models.Coverage{CoverageCode: "UninsuredMotoristPD", CvgSymbol: "1", LimitPerPerson: "100000", LimitPerOccurrence: "300000", Deductible: "500", CvgPremium: 0.0}

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
	rateresp = rtgengine.RatingEngineImpl(plcyratvars, drvratvars, vehratvars)

	return rateresp
}
