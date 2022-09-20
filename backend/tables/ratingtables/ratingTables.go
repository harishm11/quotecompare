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

type Routines struct {
	RoutineId    string `gorm:"primaryKey"`
	CoverageCode string `gorm:"primaryKey"`
	RoutineName  string
	RatebookCode string
}

type RateFactors struct {
	StateCode      string `gorm:"primaryKey"`
	ProductCode    string `gorm:"primaryKey"`
	CoverageCode   string `gorm:"primaryKey"`
	RatingItemCode string `gorm:"primaryKey"`
	RateVar1Code   string
	RateVar1Value  string
	RateVar2Code   string
	RateVar2Value  string
	RateVar3Code   string
	RateVar3Value  string
	RateVar4Code   string
	RateVar4Value  string
	RateVar5Code   string
	RateVar5Value  string
	RateVar6Code   string
	RateVar6Value  string
	RateVar7Code   string
	RateVar7Value  string
	RateVar8Code   string
	RateVar8Value  string
	RateFactor     float32
}

type RateRoutinSteps struct {
	RoutineId          string `gorm:"primaryKey"`
	CoverageCode       string `gorm:"primaryKey"`
	StepNo             uint
	RefStepNo          uint
	RatingItemCode     string
	RatingItemGrpCode  string
	StepOperation      string
	StepCalcMethod     string
	StepSplMethod      string
	CvgCodetoGetFctr   string
	RateVar1Code       string
	RateVar2Code       string
	RateVar3Code       string
	RateVar4Code       string
	RateVar5Code       string
	RateVar6Code       string
	RateVar7Code       string
	RateVar8Code       string
	DefaultValue       float32
	RoundorTrunc       string
	RoundorTruncDigits uint
}
