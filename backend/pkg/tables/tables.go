package tables

import "time"

type Ratebook struct {
	RatebookCode               string `gorm:"primaryKey"`
	RatebookActivationDate     time.Time
	RatebookNewBusinessEffDate time.Time
	RatebookRenBusinessEffDate time.Time
}
