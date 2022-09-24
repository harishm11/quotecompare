package models

type Coverage struct {
	ID                 uint `gorm:"primaryKey;uniqueIndex;"`
	VehicleID          uint `gorm:"foreignKey:ID"`
	CoverageCode       string
	CvgSymbol          string
	LimitPerPerson     string
	LimitPerOccurrence string
	Deductible         string
	CvgPremium         float32
}
