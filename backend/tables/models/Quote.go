package models

import (
	"encoding/json"
	"time"
)

type Quote struct {
	ID                uint      `gorm:"primaryKey;uniqueIndex;"`
	Quotes            Qte       `json:"quoteformFields" gorm:"foreignKey:QuoteID" `
	Drivers           []Driver  `json:"driverformFields" gorm:"foreignKey:QuoteID" `
	Vehicles          []Vehicle `json:"vehicleformFields" gorm:"foreignKey:QuoteID" `
	RateTermStartDate time.Time
	RateAppliedDate   time.Time
	QuoteStartDate    time.Time
	GoodPremium       float32
	BetterPremium     float32
	BestPremium       float32
}

type Qte struct {
	ID                uint        `gorm:"primaryKey;uniqueIndex;"`
	QuoteID           uint        `gorm:"foreignKey:ID"`
	QuoteNumber       int         `json:"quotenumber"`
	QuoteEffDate      time.Time   `json:"effdate"`
	Policyterm        json.Number `json:"policyterm"`
	AutoUmbrellaInd   string      `json:"AutoUmbrellaInd"`
	AutoHomeInd       string      `json:"AutoHomeInd"`
	AutoHomeLifeInd   string      `json:"AutoHomeLifeInd"`
	AutoLifeInd       string      `json:"AutoLifeInd"`
	AutoRenterInd     string      `json:"AutoRenterInd"`
	AutoRenterLifeInd string      `json:"AutoRenterLifeInd"`
}
