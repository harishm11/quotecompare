package models

import "encoding/json"

type Driver struct {
	ID              uint   `gorm:"primaryKey;uniqueIndex;"`
	QuoteID         uint   `gorm:"foreignKey:ID"`
	Name            string `json:"name"`
	Age             json.Number      `json:"age"`
	Experience      json.Number    `json:"experience"`
	Course          string `json:"course"`
	Incidentdate    string `json:"incidentdate"`
	Incidenttype    string `json:"incidenttype"`
	MaritalStatCode string `json:"maritalstatcode"`
}
