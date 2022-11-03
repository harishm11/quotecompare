package models

import (
	"encoding/json"
	"time"
)

type Driver struct {
	ID                uint        `gorm:"primaryKey;uniqueIndex;"`
	QuoteID           uint        `gorm:"foreignKey:ID"`
	Name              string      `json:"name"`
	Age               json.Number `json:"age"`
	Experience        json.Number `json:"experience"`
	CourseInd         string      `json:"course"`
	Incidentdate      string      `json:"incidentdate"`
	Incidenttype      string      `json:"incidenttype"`
	MaritalStatCode   string      `json:"maritalstatcode"`
	LicenseIssueDate  time.Time   `json:"licissuedt"`
	GoodStudent       string      `json:"goodstudent"`
	DateOfBirth       time.Time   `json:"dateofbirth"`
	DriverAddedDate   time.Time   `json:"drveraddeddt"`
	Occupation        string      `json:"occupation"`
	PNIInd            string      `json:"pniind"`
	RelationShiptoPNI string      `json:"relationtopni"`
}
