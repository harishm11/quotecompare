package models

import "time"

type Quote struct {
	ID           uint      `gorm:"primaryKey;uniqueIndex;"`
	QuoteNumber  int       `json:"quotenumber"`
	QuoteEffDate time.Time `json:"effdate"`
	// Lob         string    `json:"lob"`
	Drivers           []Driver  `json:"driverformFields" gorm:"foreignKey:QuoteID" `
	Vehicles          []Vehicle `json:"vehicleformFields" gorm:"foreignKey:QuoteID" `
	RateTermStartDate time.Time
	RateAppliedDate   time.Time
	QuoteStartDate    time.Time
}

type Driver struct {
	ID           uint   `gorm:"primaryKey;uniqueIndex;"`
	QuoteID      uint   `gorm:"foreignKey:ID"`
	Name         string `json:"name"`
	Age          string `json:"age"`
	Experience   string `json:"experience"`
	Course       string `json:"course"`
	Incidentdate string `json:"incidentdate"`
	Incidenttype string `json:"incidenttype"`
}
type Vehicle struct {
	ID            uint   `gorm:"primaryKey;uniqueIndex;"`
	QuoteID       uint   `gorm:"foreignKey:ID"`
	VehYear       string `json:"vehyear"`
	VehMake       string `json:"vehmake"`
	VehModel      string `json:"vehmodel"`
	AnnualMileage string `json:"annualMileage"`
	GrgZip        string `json:"grgZip"`
}
