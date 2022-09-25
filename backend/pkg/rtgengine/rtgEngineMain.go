package rtgengine

import (
	"fmt"

	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
)

func RatingEngineImpl(pv ratingvariables.PolicyRatingVars, dv []ratingvariables.DriverRatingVars, vv []ratingvariables.VehicleRatingVars) float32 {
	var plcyprm float32
	qeffdt := pv.QuoteEffDt
	qapplieddt := pv.QuoteAppliedDt

	//select ratebook using quote effective date and quote applied date
	var ratebookcode, ratebookactvtndt = RatebookSelector(qeffdt, qapplieddt)
	pv.RatebookCode = ratebookcode
	pv.RatebookActivationDate = ratebookactvtndt

	fmt.Println(ratebookcode)
	fmt.Println(ratebookactvtndt)
	//Process routinesteps

	plcyprm = ProcessRoutinesteps(pv, dv, vv)
	return plcyprm
}
