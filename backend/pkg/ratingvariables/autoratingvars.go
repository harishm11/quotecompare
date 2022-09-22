package ratingvariables

import (
	"encoding/json"
	"time"

	"github.com/harishm11/quoteCompare/tables/models"
)

type PolicyRatingVars struct {
	AutoHomeDisc       bool
	AutoHomeLifeDisc   bool
	AutoLifeDisc       bool
	AutoRenterDisc     bool
	AutoRenterLifeDisc bool
	AutoCeaDisc        bool
	Policyterm         uint
	TotalVehicles      int
	TotalDrivers       int
	MultiCarDisc       bool
	MatureDrvDisc      bool
	PersistencyDisc    bool
}

type DriverRatingVars struct {
	DrivingExpYears json.Number
	MaritalStatCode string
	GoodStdDisc     bool
	SeniorDefDisc   bool
	DriverTrainDisc bool
	GoodDriverDisc  bool
	DPSScore        uint16
}

type VehicleRatingVars struct {
	AnnualMileage     json.Number
	AltFuelInd        bool
	AntiLockBrakeDisc bool
	AntiTheftDisc     bool
	PassiveRestDisc   bool
	ESCDisc           bool
	VehicleUsage      string
	VehicleAge        int
	Zipcode           int
	Modelyear         int
	Symbol            int
	RatedDriver       models.Driver
}

func PopPolicyRatingVars(q *models.Quote) PolicyRatingVars {

	var policyRatingVars PolicyRatingVars
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

	}
	return vehicleRatingVars
}
