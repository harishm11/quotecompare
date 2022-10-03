package rtgengine

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strconv"
	"strings"
	"time"

	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/pkg/ratingvariables"
	"github.com/harishm11/quoteCompare/tables/ratingtables"
)

var plcyvar ratingvariables.PolicyRatingVars
var drvvar []ratingvariables.DriverRatingVars
var vehvar []ratingvariables.VehicleRatingVars

type RateStep struct {
	RoutineId              string
	CoverageCode           string
	StepNo                 int
	RefStepNo              int
	RatingItemCode         string
	RatingItemGrpCode      string
	StepOperation          string
	StepCalcMethod         string
	StepSplMethod          string
	CvgCodetoGetFctr       string
	RateVar1Code           string
	RateVar1Value          string
	RateVar2Code           string
	RateVar2Value          string
	RateVar3Code           string
	RateVar3Value          string
	RateVar4Code           string
	RateVar4Value          string
	RateVar5Code           string
	RateVar5Value          string
	RateVar6Code           string
	RateVar6Value          string
	RateVar7Code           string
	RateVar7Value          string
	RateVar8Code           string
	RateVar8Value          string
	DefaultValue           float32
	RoundorTrunc           string
	RoundorTruncDigits     int
	RateFactor             float32
	StepResult             float32
	RateEffDate            time.Time
	RatebookActivationDate time.Time
	RateActivationDate     time.Time
}

var RateStepTbl []RateStep

func ProcessRoutinesteps(pv ratingvariables.PolicyRatingVars, dv []ratingvariables.DriverRatingVars, vv []ratingvariables.VehicleRatingVars) float32 {

	plcyvar = pv
	drvvar = dv
	vehvar = vv

	//Read the Routinesteps

	var steps []ratingtables.RateRoutinSteps

	//Get RoutineId based on Ratebook code
	rout_id := GetRoutinId(pv.RatebookCode)

	fmt.Println(rout_id)
	//Get Routinesteps based on RoutineId
	db := database.DBConn
	db.Table("rate_routin_steps").Where("routine_id = ? ", rout_id).Order("coverage_code , step_no").Scan(&steps)

	//create working storage table with routinesteps
	RateStepTbl = make([]RateStep, len(steps))
	var tempcvg string
	var tempcvgprem float32

	//Execute the Routinesteps for each vehicle
	for vehidx := range vehvar {
		for stpidx := range steps {

			//copy routinesteps to working storage table
			CopyRoutineStp2Tbl(stpidx, steps[stpidx])

			//populate rating variable valies in working storage table
			GetStpRatingVarValue(stpidx, vehidx)

			//populate ratebook activation date in WS table
			RateStepTbl[stpidx].RatebookActivationDate = pv.RatebookActivationDate
			RateStepTbl[stpidx].RateEffDate = pv.QuoteEffDt

			//retrieve factor and rateactivation date and store in working storage table
			RateStepTbl[stpidx].RateFactor, RateStepTbl[stpidx].RateActivationDate = GetRatingFactor(RateStepTbl[stpidx])

			if RateStepTbl[stpidx].CoverageCode != tempcvg {
				RateStepTbl[stpidx].StepResult = RateStepTbl[stpidx].RateFactor
				tempcvg = RateStepTbl[stpidx].CoverageCode
			} else {
				RateStepTbl[stpidx].StepResult = RateStepTbl[stpidx-1].StepResult * RateStepTbl[stpidx].RateFactor
			}

			if RateStepTbl[stpidx].StepCalcMethod == "StoreResult" {
				for cvgidx := range vehvar[vehidx].CoverageRatingVars {
					if RateStepTbl[stpidx].CoverageCode == vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode {
						tempcvgprem = RateStepTbl[stpidx].StepResult
						vehvar[vehidx].CoverageRatingVars[cvgidx].CvgPremium = RateStepTbl[stpidx].StepResult
						vehvar[vehidx].VehPremium = vehvar[vehidx].VehPremium + tempcvgprem
						plcyvar.PlcyPremium = plcyvar.PlcyPremium + tempcvgprem
						fmt.Println(tempcvg, "Coverage premium = ", tempcvgprem)
					}
				}
			}
		}
		fmt.Println("Veh", vehidx+1, "premium = ", vehvar[vehidx].VehPremium)
	}
	fmt.Println("Policy Premium = ", plcyvar.PlcyPremium)
	//fmt.Println(len(RateStepTbl))
	Generateworksheet(RateStepTbl)
	return plcyvar.PlcyPremium

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

	case "AffinityGroup":
		RateVarValue = plcyvar.AffinityGroup
	case "HouseholdCompostion":
		RateVarValue = plcyvar.HouseholdCompostion
	case "TermLength":
		RateVarValue = strconv.Itoa(plcyvar.TermLength)
	case "MultiPolicy":
		RateVarValue = plcyvar.MultiPolicy
	case "PermissiveUserOption":
		RateVarValue = plcyvar.PermissiveUserOption

	case "Zipcode":
		RateVarValue = vehvar[vehidx].Zipcode
	case "Mileage1":
		RateVarValue = vehvar[vehidx].Mileage1.String()
	case "Mileage2":
		RateVarValue = vehvar[vehidx].Mileage2.String()
	case "VehicleAge1":
		RateVarValue = strconv.Itoa(vehvar[vehidx].VehicleAge1)
	case "VehicleAge2":
		RateVarValue = strconv.Itoa(vehvar[vehidx].VehicleAge2)
	case "ModelYear":
		RateVarValue = strconv.Itoa(vehvar[vehidx].ModelYear)
	case "HighPerfInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].HighPerfInd))
	case "VehHistInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].VehHistInd))
	case "PassiveResType":
		RateVarValue = vehvar[vehidx].PassiveResType
	case "AntiLockInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].AntiLockInd))
	case "AntitheftInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].AntitheftInd))
	case "AltFuelInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].AltFuelInd))
	case "EscInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].EscInd))
	case "VehUseCode":
		RateVarValue = vehvar[vehidx].VehUseCode
	case "FrequencyBand":
		RateVarValue = vehvar[vehidx].FrequencyBand
	case "SeverityBand":
		RateVarValue = vehvar[vehidx].SeverityBand
	case "RideShareInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(vehvar[vehidx].RideShareInd))

	case "DPS":
		RateVarValue = strconv.Itoa(drvvar[0].DPS)
	case "DriverClass":
		RateVarValue = strconv.Itoa(drvvar[0].DriverClass)
	case "DriverClassCode":
		RateVarValue = drvvar[0].DriverClassCode
	case "YearsDrivingExperience":
		RateVarValue = "All Other"
		//drvvar[0].YearsDrivingExperience.String()
	case "StudentAwayInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(drvvar[0].StudentAwayInd))
	case "MultiCarDisc":
		RateVarValue = strings.ToUpper(strconv.FormatBool(plcyvar.MultiCarDisc))
	case "MatureDrvDisc":
		RateVarValue = strings.ToUpper(strconv.FormatBool(plcyvar.MatureDrvDisc))
	case "PersistencyDisc":
		RateVarValue = strings.ToUpper(strconv.FormatBool(plcyvar.PersistencyDisc))
	case "GoodDriverDiscInd":
		RateVarValue = strings.ToUpper(strconv.FormatBool(drvvar[0].GoodDriverDiscInd))
	case "MaritalStatCode":
		RateVarValue = drvvar[0].MaritalStatCode

	case "Limit1":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].Limit1
			}
		}
	case "Limit2":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].Limit2
			}
		}
	case "Symbol":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].CvgSymbol
			}

		}
	case "Deductible1":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].Deductible1
			}

		}
	case "Deductible2":
		for cvgidx := range vehvar[vehidx].CoverageRatingVars {
			if vehvar[vehidx].CoverageRatingVars[cvgidx].CoverageCode == RateStepTbl[stpidx].CoverageCode {
				RateVarValue = vehvar[vehidx].CoverageRatingVars[cvgidx].Deductible2
			}

		}

	}
	return RateVarValue
}

func Generateworksheet(RateStepTbl []RateStep) {

	dir, err := ioutil.ReadDir("download")
	if err != nil {
		fmt.Println(err)
	}
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"download", d.Name()}...))
	}

	file, _ := json.MarshalIndent(RateStepTbl, "", " ")
	_ = ioutil.WriteFile("download/output.json", file, 0644)

	data, err := ioutil.ReadFile("download/output.json")
	if err != nil {
		fmt.Println(err)
	}
	// Unmarshal JSON data
	var d []RateStep
	err = json.Unmarshal([]byte(data), &d)
	if err != nil {
		fmt.Println(err)
	}
	// Create a csv file
	f, err := os.Create("download/output.csv")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()
	// Write Unmarshaled json data to CSV file
	w := csv.NewWriter(f)
	for _, obj := range d {
		var record []string
		record = append(
			record,
			obj.RoutineId,
			obj.CoverageCode,
			strconv.Itoa(obj.StepNo),
			strconv.Itoa(obj.RefStepNo),
			obj.RatingItemCode,
			obj.RatingItemGrpCode,
			obj.StepOperation,
			obj.StepCalcMethod,
			obj.StepSplMethod,
			obj.CvgCodetoGetFctr,
			obj.RateVar1Code,
			obj.RateVar1Value,
			obj.RateVar2Code,
			obj.RateVar2Value,
			obj.RateVar3Code,
			obj.RateVar3Value,
			obj.RateVar4Code,
			obj.RateVar4Value,
			obj.RateVar5Code,
			obj.RateVar5Value,
			obj.RateVar6Code,
			obj.RateVar6Value,
			obj.RateVar7Code,
			obj.RateVar7Value,
			obj.RateVar8Code,
			obj.RateVar8Value,
			fmt.Sprintf("%f", obj.DefaultValue),
			obj.RoundorTrunc,
			strconv.Itoa(obj.RoundorTruncDigits),
			fmt.Sprintf("%f", obj.RateFactor),
			fmt.Sprintf("%f", obj.StepResult),
			obj.RateEffDate.String(),
			obj.RatebookActivationDate.String(),
			obj.RateActivationDate.String())
		w.Write(record)
	}
	w.Flush()

}
