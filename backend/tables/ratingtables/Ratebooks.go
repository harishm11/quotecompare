package ratingtables

import "time"

type Ratebooks struct {
	StateCode                  string `gorm:"primaryKey"`
	ProductCode                string `gorm:"primaryKey"`
	RatebookCode               string `gorm:"primaryKey"`
	RatebookActivationDate     time.Time
	RatebookNewBusinessEffDate time.Time
	RatebookRenBusinessEffDate time.Time
}