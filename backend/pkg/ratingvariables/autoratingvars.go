package ratingvariables

import (
	"github.com/harishm11/quoteCompare/tables/models"
)

type PolicyRatingVars struct {
	AutoHomeDisc       bool
	AutoHomeLifeDisc   bool
	AutoLifeDisc       bool
	AutoRenterDisc     bool
	AutoRenterLifeDisc bool
	AutoCeaDisc        bool
	Policyterm         int
}

type DriverRatingVars struct {
	DrivingExpYears string
	MaritalStatCode string
	GoodStdDisc     bool
	SeniorDefDisc   bool
	DriverTrainDisc bool
	GoodDriverDisc  bool
	TotalDrivers    uint16
	DPSScore        uint16
}

type VehicleRatingVars struct {
	AnnualMileage     string
	AltFuelInd        bool
	AntiLockBrakeDisc bool
	AntiTheftDisc     bool
	TotalVehicles     uint16
	PassiveRestDisc   bool
	ESCDisc           bool
	VehicleUsage      string
	VehicleAge        int
}

func PopPolicyRatingVars(q *models.Quote) PolicyRatingVars {

	var policyRatingVars PolicyRatingVars

	policyRatingVars.AutoHomeDisc = false
	policyRatingVars.AutoHomeLifeDisc = false
	policyRatingVars.AutoLifeDisc = false
	policyRatingVars.AutoRenterDisc = false
	policyRatingVars.AutoRenterLifeDisc = false
	policyRatingVars.AutoCeaDisc = false
	policyRatingVars.Policyterm = q.Policyterm

	return policyRatingVars
}

func PopDriverRatingVars(d []models.Driver) []DriverRatingVars {
	var driverRatingVars = make([]DriverRatingVars, len(d))

	for index, drv := range d {
		driverRatingVars[index].GoodStdDisc = false
		driverRatingVars[index].GoodDriverDisc = false
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
		vehicleRatingVars[index].VehicleAge = ' '
		//time.Now().Year() - veh.VehYear
		vehicleRatingVars[index].VehicleUsage = veh.VehicleUsage

	}
	return vehicleRatingVars
}
