package models

type Driver struct {
	ID              uint   `gorm:"primaryKey;uniqueIndex;"`
	QuoteID         uint   `gorm:"foreignKey:ID"`
	Name            string `json:"name"`
	Age             string `json:"age"`
	Experience      string `json:"experience"`
	Course          string `json:"course"`
	Incidentdate    string `json:"incidentdate"`
	Incidenttype    string `json:"incidenttype"`
	MaritalStatCode string `json:"maritalstatcode"`
}
