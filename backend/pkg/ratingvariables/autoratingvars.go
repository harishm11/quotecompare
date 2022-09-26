package ratingvariables

import (
	"encoding/json"
	"time"

	"github.com/harishm11/quoteCompare/tables/models"
)

type PolicyRatingVars struct {
	QuoteEffDt             time.Time
	QuoteAppliedDt         time.Time
	AutoHomeDisc           bool
	AutoHomeLifeDisc       bool
	AutoLifeDisc           bool
	AutoRenterDisc         bool
	AutoRenterLifeDisc     bool
	AutoCeaDisc            bool
	Policyterm             int
	TotalVehicles          int
	TotalDrivers           int
	MultiCarDisc           bool
	MatureDrvDisc          bool
	PersistencyDisc        bool
	RatebookCode           string
	RatebookActivationDate time.Time
	PlcyPremium            float32
}

type DriverRatingVars struct {
	DriverId        uint
	DrivingExpYears json.Number
	MaritalStatCode string
	GoodStdDisc     bool
	SeniorDefDisc   bool
	DriverTrainDisc bool
	GoodDriverDisc  bool
	Points          int
}

type VehicleRatingVars struct {
	VehicleId          uint
	AnnualMileage      json.Number
	AltFuelInd         bool
	AntiLockBrakeDisc  bool
	AntiTheftDisc      bool
	PassiveRestDisc    bool
	ESCDisc            bool
	VehicleUsage       string
	VehicleAge         int
	Zipcode            string
	Modelyear          int
	Symbol             int
	VehPremium         float32
	CoverageRatingVars []CoverageRatingVars
}

type CoverageRatingVars struct {
	CoverageCode       string
	CvgSymbol          string
	LimitPerPerson     string
	LimitPerOccurrence string
	Deductible         string
	CvgPremium         float32
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
	policyRatingVars.AutoHomeDisc = false
	policyRatingVars.AutoHomeLifeDisc = false
	policyRatingVars.AutoLifeDisc = false
	policyRatingVars.AutoRenterDisc = false
	policyRatingVars.AutoRenterLifeDisc = false
	policyRatingVars.AutoCeaDisc = false
	policyRatingVars.PersistencyDisc = false
	policyRatingVars.Policyterm = q.Policyterm

	return policyRatingVars
}

func PopDriverRatingVars(d []models.Driver) []DriverRatingVars {
	var driverRatingVars = make([]DriverRatingVars, len(d))

	for index, drv := range d {
		driverRatingVars[index].DriverId = drv.ID
		driverRatingVars[index].GoodStdDisc = false
		driverRatingVars[index].GoodDriverDisc = true
		driverRatingVars[index].DriverTrainDisc = false
		driverRatingVars[index].SeniorDefDisc = false
		driverRatingVars[index].MaritalStatCode = drv.MaritalStatCode
		driverRatingVars[index].DrivingExpYears = drv.Experience
	}
	return driverRatingVars
}

func PopVehicleRatingVars(v []models.Vehicle) []VehicleRatingVars {
	var vehicleRatingVars = make([]VehicleRatingVars, len(v))

	for index, veh := range v {
		vehicleRatingVars[index].VehicleId = veh.ID
		vehicleRatingVars[index].AnnualMileage = veh.AnnualMileage
		vehicleRatingVars[index].AntiLockBrakeDisc = false
		vehicleRatingVars[index].AntiTheftDisc = false
		vehicleRatingVars[index].PassiveRestDisc = false
		vehicleRatingVars[index].PassiveRestDisc = false
		vehicleRatingVars[index].ESCDisc = false
		var err error
		var value int64
		value, err = veh.VehYear.Int64()
		if err != nil {
			vehicleRatingVars[index].VehicleAge = time.Now().Year() - int(value)
		}
		vehicleRatingVars[index].Modelyear = int(value)
		vehicleRatingVars[index].VehicleUsage = veh.VehicleUsage
		vehicleRatingVars[index].Symbol = 1
		vehicleRatingVars[index].Zipcode = veh.GrgZip
		vehicleRatingVars[index].CoverageRatingVars = PopCvgRatingVars(veh.Coverages)

	}
	return vehicleRatingVars
}

func PopCvgRatingVars(c []models.Coverage) []CoverageRatingVars {
	var coverageRatingVars = make([]CoverageRatingVars, len(c))
	for index, cvg := range c {
		coverageRatingVars[index].CoverageCode = cvg.CoverageCode
		coverageRatingVars[index].CvgSymbol = cvg.CvgSymbol
		coverageRatingVars[index].LimitPerPerson = cvg.LimitPerPerson
		coverageRatingVars[index].LimitPerOccurrence = cvg.LimitPerOccurrence
		coverageRatingVars[index].Deductible = cvg.Deductible
		coverageRatingVars[index].CvgPremium = cvg.CvgPremium
	}
	return coverageRatingVars
}
