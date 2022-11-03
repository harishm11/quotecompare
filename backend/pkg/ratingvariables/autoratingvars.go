package ratingvariables

import (
	"encoding/json"
	"time"

	"github.com/harishm11/quoteCompare/tables/models"
)

type PolicyRatingVars struct {
	QuoteEffDt             time.Time
	QuoteAppliedDt         time.Time
	MultiPolicy            string
	TermLength             int
	TotalVehicles          int
	TotalDrivers           int
	MultiCarDisc           bool
	MatureDrvDisc          bool
	PersistencyDisc        bool
	RatebookCode           string
	RatebookActivationDate time.Time
	PlcyPremium            float32
	HouseholdCompostion    string
	PermissiveUserOption   string
	AffinityGroup          string
}

type DriverRatingVars struct {
	DriverId               uint
	YearsDrivingExperience json.Number
	MaritalStatCode        string
	GoodStdDisc            bool
	SeniorDefDisc          bool
	DriverTrainDisc        bool
	GoodDriverDiscInd      bool
	DPS                    int
	DriverClass            int
	DriverClassCode        string
	StudentAwayInd         bool
}

type VehicleRatingVars struct {
	VehicleId          uint
	Mileage1           json.Number
	Mileage2           json.Number
	AltFuelInd         bool
	AntiLockInd        bool
	AntitheftInd       bool
	PassiveResType     string
	EscInd             bool
	HighPerfInd        bool
	VehHistInd         bool
	VehUseCode         string
	VehicleAge1        int
	VehicleAge2        int
	Zipcode            string
	ModelYear          int
	Symbol             int
	VehPremium         float32
	FrequencyBand      string
	SeverityBand       string
	RideShareInd       bool
	CoverageRatingVars []CoverageRatingVars
}

type CoverageRatingVars struct {
	CoverageCode string
	CvgSymbol    string
	Limit1       string
	Limit2       string
	Deductible1  string
	Deductible2  string
	CvgPremium   float32
}

func PopPolicyRatingVars(q *models.Quote) PolicyRatingVars {

	var policyRatingVars PolicyRatingVars
	policyRatingVars.QuoteEffDt = q.QuoteEffDate
	policyRatingVars.QuoteAppliedDt = q.RateAppliedDate
	policyRatingVars.TotalVehicles = len(q.Vehicles)
	policyRatingVars.TotalDrivers = len(q.Drivers)
	if policyRatingVars.TotalVehicles > 1 {
		policyRatingVars.MultiCarDisc = true
	}
	policyRatingVars.MatureDrvDisc = false
	policyRatingVars.MultiPolicy = "Auto-Life"
	policyRatingVars.PersistencyDisc = false
	policyRatingVars.TermLength = q.Policyterm
	policyRatingVars.HouseholdCompostion = "DV-11"
	policyRatingVars.PermissiveUserOption = "Full"
	policyRatingVars.AffinityGroup = "Group I"

	return policyRatingVars
}

func PopDriverRatingVars(d []models.Driver) []DriverRatingVars {
	var driverRatingVars = make([]DriverRatingVars, len(d))

	for index, drv := range d {
		driverRatingVars[index].DriverId = drv.ID
		driverRatingVars[index].GoodStdDisc = false
		driverRatingVars[index].GoodDriverDiscInd = true
		driverRatingVars[index].DriverTrainDisc = false
		driverRatingVars[index].SeniorDefDisc = false
		driverRatingVars[index].MaritalStatCode = drv.MaritalStatCode
		driverRatingVars[index].YearsDrivingExperience = drv.Experience
		driverRatingVars[index].DriverClass = 1
		driverRatingVars[index].DriverClassCode = "0FMN"
		driverRatingVars[index].StudentAwayInd = false
		driverRatingVars[index].DPS = 0
	}
	return driverRatingVars
}

func PopVehicleRatingVars(v []models.Vehicle) []VehicleRatingVars {
	var vehicleRatingVars = make([]VehicleRatingVars, len(v))

	for index, veh := range v {
		vehicleRatingVars[index].VehicleId = veh.ID
		vehicleRatingVars[index].Mileage1 = veh.AnnualMileage
		vehicleRatingVars[index].Mileage2 = veh.AnnualMileage
		vehicleRatingVars[index].AntiLockInd = false
		vehicleRatingVars[index].AntitheftInd = false
		vehicleRatingVars[index].PassiveResType = "BELTS"
		vehicleRatingVars[index].EscInd = false
		vehicleRatingVars[index].HighPerfInd = false
		vehicleRatingVars[index].VehHistInd = false
		var err error
		var value int64
		value, err = veh.VehYear.Int64()
		if err != nil {
			vehicleRatingVars[index].VehicleAge1 = time.Now().Year() - int(value)
			vehicleRatingVars[index].VehicleAge2 = vehicleRatingVars[index].VehicleAge1
		}
		vehicleRatingVars[index].ModelYear = int(value)
		vehicleRatingVars[index].VehUseCode = "Business"
		//veh.VehicleUsage
		vehicleRatingVars[index].Symbol = 1
		vehicleRatingVars[index].Zipcode = veh.GrgZip
		vehicleRatingVars[index].FrequencyBand = "1"
		vehicleRatingVars[index].SeverityBand = "2"
		vehicleRatingVars[index].RideShareInd = false
		vehicleRatingVars[index].CoverageRatingVars = PopCvgRatingVars(veh.Coverages)

	}
	return vehicleRatingVars
}

func PopCvgRatingVars(c []models.Coverage) []CoverageRatingVars {
	var coverageRatingVars = make([]CoverageRatingVars, len(c))
	for index, cvg := range c {
		coverageRatingVars[index].CoverageCode = cvg.CoverageCode
		coverageRatingVars[index].CvgSymbol = cvg.CvgSymbol
		coverageRatingVars[index].Limit1 = cvg.LimitPerPerson
		coverageRatingVars[index].Limit2 = cvg.LimitPerOccurrence
		coverageRatingVars[index].Deductible1 = cvg.Deductible
		coverageRatingVars[index].Deductible2 = ""
		coverageRatingVars[index].CvgPremium = cvg.CvgPremium
	}
	return coverageRatingVars
}
