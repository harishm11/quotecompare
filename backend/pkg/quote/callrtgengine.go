package quote

import (
	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
	"github.com/harishm11/quoteCompare/pkg/rtgengine"
	"github.com/harishm11/quoteCompare/tables/models"
)

func InvokeRtgEngine(quote *models.Quote) {

	var c1 = models.Coverage{CoverageCode: "BodilyInjury", CvgSymbol: "", LimitPerPerson: "100000", LimitPerOccurrence: "300000", Deductible: "", CvgPremium: 0.0}
	var c2 = models.Coverage{CoverageCode: "PropertyDamage", CvgSymbol: "", LimitPerPerson: "100000", LimitPerOccurrence: "", Deductible: "", CvgPremium: 0.0}
	var c3 = models.Coverage{CoverageCode: "Comprehensive", CvgSymbol: "", LimitPerPerson: "", LimitPerOccurrence: "", Deductible: "500", CvgPremium: 0.0}
	var c4 = models.Coverage{CoverageCode: "Medical", CvgSymbol: "", LimitPerPerson: "5000", LimitPerOccurrence: "", Deductible: "", CvgPremium: 0.0}
	var c5 = models.Coverage{CoverageCode: "Collission", CvgSymbol: "", LimitPerPerson: "", LimitPerOccurrence: "", Deductible: "500", CvgPremium: 0.0}
	var c6 = models.Coverage{CoverageCode: "UninsuredMotoristBI", CvgSymbol: "", LimitPerPerson: "100000", LimitPerOccurrence: "300000", Deductible: "", CvgPremium: 0.0}
	var c7 = models.Coverage{CoverageCode: "UninsuredMotoristPD", CvgSymbol: "", LimitPerPerson: "100000", LimitPerOccurrence: "300000", Deductible: "", CvgPremium: 0.0}

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
	quote.GoodPremium = rtgengine.RatingEngineImpl(plcyratvars, drvratvars, vehratvars)
	// quote.BetterPremium = rtgengine.RatingEngineImpl(plcyratvars, drvratvars, vehratvars)
	// quote.BestPremium = rtgengine.RatingEngineImpl(plcyratvars, drvratvars, vehratvars)

}
