package ratingtables

type Routines struct {
	RoutineId    string `gorm:"primaryKey"`
	CoverageCode string `gorm:"primaryKey"`
	RoutineName  string
	RatebookCode string
}
