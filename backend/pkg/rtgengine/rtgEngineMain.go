package rtgengine

import (
	"fmt"

	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
)

func RatingEngineImpl(pv ratingvariables.PolicyRatingVars, dv []ratingvariables.DriverRatingVars, vv []ratingvariables.VehicleRatingVars) {

	qeffdt := pv.QuoteEffDt
	qapplieddt := pv.QuoteAppliedDt

	//select ratebook using quote effective date and quote applied date
	var ratebookcode, ratebookactvtndt = RatebookSelector(qeffdt, qapplieddt)
	pv.RatebookCode = ratebookcode
	pv.RatebookActivationDate = ratebookactvtndt

	fmt.Println(ratebookcode)
	fmt.Println(ratebookactvtndt)
	//Process routinesteps

	ProcessRoutinesteps(pv, dv, vv)

}
