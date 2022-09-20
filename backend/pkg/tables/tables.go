package tables

import "time"

type Ratebook struct {
	ProductCode                string `gorm:"primaryKey"`
	RatebookCode               string `gorm:"primaryKey"`
	RatebookActivationDate     time.Time
	RatebookNewBusinessEffDate time.Time
	RatebookRenBusinessEffDate time.Time
}

type Routine struct {
	RoutineID    string `gorm:"primaryKey"`
	CoverageCode string
	RoutineName  string
	RatebookCode string
}
