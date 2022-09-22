package models

import "encoding/json"

type Vehicle struct {
	ID            uint        `gorm:"primaryKey;uniqueIndex;"`
	QuoteID       uint        `gorm:"foreignKey:ID"`
	VehYear       json.Number `json:"vehyear"`
	VehMake       string      `json:"vehmake"`
	VehModel      string      `json:"vehmodel"`
	AnnualMileage json.Number `json:"annualMileage"`
	GrgZip        string      `json:"grgZip"`
	VehicleUsage  string      `json:"vehicleusage"`
}
