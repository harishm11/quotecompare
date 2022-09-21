package models

type Vehicle struct {
	ID            uint   `gorm:"primaryKey;uniqueIndex;"`
	QuoteID       uint   `gorm:"foreignKey:ID"`
	VehYear       string `json:"vehyear"`
	VehMake       string `json:"vehmake"`
	VehModel      string `json:"vehmodel"`
	AnnualMileage string `json:"annualMileage"`
	GrgZip        string `json:"grgZip"`
	VehicleUsage  string `json:"vehicleusage"`
}
