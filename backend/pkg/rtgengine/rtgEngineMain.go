package rtgengine

import (
	"github.com/harishm11/quoteCompare/contracts"
	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
)

func RatingEngineImpl(pv ratingvariables.PolicyRatingVars, dv []ratingvariables.DriverRatingVars, vv []ratingvariables.VehicleRatingVars) *contracts.RateResponse {
	qeffdt := pv.QuoteEffDt
	qapplieddt := pv.QuoteAppliedDt

	//select ratebook using quote effective date and quote applied date
	var ratebookcode, ratebookactvtndt = RatebookSelector(qeffdt, qapplieddt)
	pv.RatebookCode = ratebookcode
	pv.RatebookActivationDate = ratebookactvtndt

	//Process routinesteps
	rateresp := ProcessRoutinesteps(pv, dv, vv)
	return rateresp
}
