package rtgengine

import (
	"fmt"
	"strconv"

	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
	"github.com/harishm11/quoteCompare/tables/ratingtables"
)

var plcyvar ratingvariables.PolicyRatingVars
var drvvar []ratingvariables.DriverRatingVars
var vehvar []ratingvariables.VehicleRatingVars

type RateStep struct {
	RoutineId          string
	CoverageCode       string
	StepNo             int
	RefStepNo          int
	RatingItemCode     string
	RatingItemGrpCode  string
	StepOperation      string
	StepCalcMethod     string
	StepSplMethod      string
	CvgCodetoGetFctr   string
	RateVar1Code       string
	RateVar1Value      string
	RateVar2Code       string
	RateVar2Value      string
	RateVar3Code       string
	RateVar3Value      string
	RateVar4Code       string
	RateVar4Value      string
	RateVar5Code       string
	RateVar5Value      string
	RateVar6Code       string
	RateVar6Value      string
	RateVar7Code       string
	RateVar7Value      string
	RateVar8Code       string
	RateVar8Value      string
	DefaultValue       float32
	RoundorTrunc       string
	RoundorTruncDigits int
	RateFactor         float32
}

var RateStepTbl []RateStep

func ProcessRoutinesteps(pv ratingvariables.PolicyRatingVars, dv []ratingvariables.DriverRatingVars, vv []ratingvariables.VehicleRatingVars) {

	plcyvar = pv
	drvvar = dv
	vehvar = vv

	//Read the Routinesteps

	var steps []ratingtables.RateRoutinSteps
	rout_id := "R001"
	db := database.DBConn
	db.Table("rate_routin_steps").Where("routine_id = ? ", rout_id).Scan(&steps)

	RateStepTbl = make([]RateStep, len(steps))

	//Execute the Routinesteps for each vehicle

	for vehidx := range vehvar {
		for stpidx := range steps {
			CopyRoutineStp2Tbl(stpidx, steps[stpidx])
			GetStpRatingVarValue(stpidx, vehidx)
			fmt.Println(RateStepTbl[stpidx])
			// if RateStepTbl[stpidx].RatingItemCode == "RI001" || RateStepTbl[stpidx].RatingItemCode == "RI002" {
			// 	GetRatingFactor(stpidx)
			// }

		}

	}

}

func CopyRoutineStp2Tbl(i int, s ratingtables.RateRoutinSteps) {
	RateStepTbl[i].RoutineId = s.RoutineId
	RateStepTbl[i].CoverageCode = s.CoverageCode
	RateStepTbl[i].StepNo = s.StepNo
	RateStepTbl[i].RefStepNo = s.RefStepNo
	RateStepTbl[i].RatingItemCode = s.RatingItemCode
	RateStepTbl[i].RatingItemGrpCode = s.RatingItemGrpCode
	RateStepTbl[i].StepOperation = s.StepOperation
	RateStepTbl[i].StepCalcMethod = s.StepCalcMethod
	RateStepTbl[i].StepSplMethod = s.StepSplMethod
	RateStepTbl[i].CvgCodetoGetFctr = s.CvgCodetoGetFctr
	RateStepTbl[i].RateVar1Code = s.RateVar1Code
	RateStepTbl[i].RateVar2Code = s.RateVar2Code
	RateStepTbl[i].RateVar3Code = s.RateVar3Code
	RateStepTbl[i].RateVar4Code = s.RateVar4Code
	RateStepTbl[i].RateVar5Code = s.RateVar5Code
	RateStepTbl[i].RateVar6Code = s.RateVar6Code
	RateStepTbl[i].RateVar7Code = s.RateVar7Code
	RateStepTbl[i].RateVar8Code = s.RateVar8Code
	RateStepTbl[i].DefaultValue = s.DefaultValue
	RateStepTbl[i].RoundorTrunc = s.RoundorTrunc
	RateStepTbl[i].RoundorTruncDigits = s.RoundorTruncDigits
}

func GetStpRatingVarValue(stpidx int, vehidx int) {
	//lookup rating variable value
	var RateVarCode string
	if RateStepTbl[stpidx].RateVar1Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar1Code
		RateStepTbl[stpidx].RateVar1Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar2Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar2Code
		RateStepTbl[stpidx].RateVar2Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar3Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar3Code
		RateStepTbl[stpidx].RateVar3Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar4Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar4Code
		RateStepTbl[stpidx].RateVar4Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar5Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar5Code
		RateStepTbl[stpidx].RateVar5Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar6Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar6Code
		RateStepTbl[stpidx].RateVar6Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar7Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar7Code
		RateStepTbl[stpidx].RateVar7Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}
	if RateStepTbl[stpidx].RateVar8Code != "" {
		RateVarCode = RateStepTbl[stpidx].RateVar8Code
		RateStepTbl[stpidx].RateVar8Value = LookupRatVarValue(stpidx, vehidx, RateVarCode)
	}

}

func LookupRatVarValue(stpidx int, vehidx int, RateVarCode string) string {
	var RateVarValue string
	switch RateVarCode {
	case "Zipcode":
		RateVarValue = vehvar[vehidx].Zipcode
	case "Points":
		RateVarValue = strconv.Itoa(drvvar[0].Points)
	case "AnnualMileage":
		RateVarValue = vehvar[vehidx].AnnualMileage.String()
	case "DrivingExpYears":
		RateVarValue = drvvar[0].DrivingExpYears.String()
	case "MultiCarDisc":
		RateVarValue = strconv.FormatBool(plcyvar.MultiCarDisc)
	case "MatureDrvDisc":
		RateVarValue = strconv.FormatBool(plcyvar.MatureDrvDisc)
	case "PersistencyDisc":
		RateVarValue = strconv.FormatBool(plcyvar.PersistencyDisc)
	case "GoodDriverDisc":
		RateVarValue = strconv.FormatBool(drvvar[0].GoodDriverDisc)
	case "MaritalStatCode":
		RateVarValue = drvvar[0].MaritalStatCode
	case "PolicyTerm":
		RateVarValue = strconv.Itoa(plcyvar.Policyterm)
	case "LimitPerPerson":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].LimitPerPerson
			}
		}
	case "LimitPerOccurrence":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].LimitPerOccurrence
			}

		}
	case "Deductible":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].Deductible
			}

		}

	}
	return RateVarValue
}

func GetRatingFactor(stpidx int) {
	var factorRow ratingtables.RateFactors
	db := database.DBConn
	db.Table("rate_factors").Where("coverage_code = ? AND rating_item_code = ? AND (rate_var1_code= ? OR rate_var1_code is NULL ) AND	(rate_var1_value =? OR rate_var1_value is NULL ) AND (rate_var2_code= ? OR rate_var2_code is NULL )AND (rate_var2_value =? OR rate_var2_value is NULL )AND (rate_var3_code= ? OR rate_var3_code is NULL ) AND (rate_var3_value =? OR rate_var3_value is NULL )AND (rate_var4_code= ? OR rate_var4_code is NULL )AND (rate_var4_value =? OR rate_var4_value is NULL )AND(rate_var5_code= ? OR rate_var5_code is NULL ) AND	(rate_var5_value =? OR rate_var5_value is NULL ) AND (rate_var6_code= ? OR rate_var6_code is NULL )AND (rate_var6_value =? OR rate_var6_value is NULL )AND (rate_var7_code= ? OR rate_var7_code is NULL ) AND (rate_var7_value =? OR rate_var7_value is NULL )AND (rate_var8_code= ? OR rate_var8_code is NULL )AND (rate_var8_value =? OR rate_var8_value is NULL )  ",
		RateStepTbl[stpidx].CvgCodetoGetFctr, RateStepTbl[stpidx].RatingItemCode,
		RateStepTbl[stpidx].RateVar1Code, RateStepTbl[stpidx].RateVar1Value,
		RateStepTbl[stpidx].RateVar2Code, RateStepTbl[stpidx].RateVar2Value,
		RateStepTbl[stpidx].RateVar3Code, RateStepTbl[stpidx].RateVar3Value,
		RateStepTbl[stpidx].RateVar4Code, RateStepTbl[stpidx].RateVar4Value,
		RateStepTbl[stpidx].RateVar5Code, RateStepTbl[stpidx].RateVar5Value,
		RateStepTbl[stpidx].RateVar6Code, RateStepTbl[stpidx].RateVar6Value,
		RateStepTbl[stpidx].RateVar7Code, RateStepTbl[stpidx].RateVar7Value,
		RateStepTbl[stpidx].RateVar8Code, RateStepTbl[stpidx].RateVar8Value).Select("rate_factor").First(&factorRow)

	// if err != nil {
	// 	RateStepTbl[stpidx].RateFactor = 1.00
	// 	fmt.Println(err)
	// } else {

	// }
	RateStepTbl[stpidx].RateFactor = factorRow.RateFactor
	fmt.Println(RateStepTbl[stpidx])

}
