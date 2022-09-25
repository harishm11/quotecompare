package models

import "time"

type Quote struct {
	ID           uint      `gorm:"primaryKey;uniqueIndex;"`
	QuoteNumber  int       `json:"quotenumber"`
	QuoteEffDate time.Time `json:"effdate"`
	// Lob         string    `json:"lob"`
	Drivers                []Driver  `json:"driverformFields" gorm:"foreignKey:QuoteID" `
	Vehicles               []Vehicle `json:"vehicleformFields" gorm:"foreignKey:QuoteID" `
	RateTermStartDate      time.Time
	RateAppliedDate        time.Time
	QuoteStartDate         time.Time
	Policyterm             int
	
}
