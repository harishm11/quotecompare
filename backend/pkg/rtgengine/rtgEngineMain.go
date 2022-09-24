package rtgengine

import (
	"fmt"

	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
)

func RatingEngineImpl(pv ratingvariables.PolicyRatingVars, dv []ratingvariables.DriverRatingVars, vv []ratingvariables.VehicleRatingVars) {

	qeffdt := pv.QuoteEffDt
	qapplieddt := pv.QuoteAppliedDt

	//select ratebook using quote effective date and quote applied date
	var ratebookcode = RatebookSelector(qeffdt, qapplieddt)
	fmt.Println("Quote eff date ", pv.QuoteEffDt)
	fmt.Println("Ratebook code", ratebookcode)

	//Process routinesteps
	ProcessRoutinesteps(pv, dv, vv)

}
