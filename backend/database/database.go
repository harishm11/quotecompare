package database

import (
	"fmt"

	"github.com/harishm11/quoteCompare/tables/models"
	"github.com/harishm11/quoteCompare/tables/ratingtables"
	"github.com/harishm11/quoteCompare/tables/ruletables"
	"github.com/harishm11/quoteCompare/tables/systemtables"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DBConn *gorm.DB
)

func InitDatabase() {
	var err error
	// DBConn, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	dsn := "host=localhost user=postgres password=readonly dbname=postgres port=5432 sslmode=disable"
	DBConn, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	fmt.Println("Database connection successfully opened")

	//Policy Tables
	DBConn.AutoMigrate(&models.Quote{})
	DBConn.AutoMigrate(&models.Driver{})
	DBConn.AutoMigrate(&models.Vehicle{})
	DBConn.AutoMigrate(&models.Coverage{})

	//Rating tables
	DBConn.AutoMigrate(&ratingtables.Ratebooks{})
	DBConn.AutoMigrate(&ratingtables.RateRoutines{})
	DBConn.AutoMigrate(&ratingtables.RateFactors{})
	DBConn.AutoMigrate(&ratingtables.RateRoutinSteps{})

	//System tables
	DBConn.AutoMigrate(&systemtables.States{})
	DBConn.AutoMigrate(&systemtables.CompanyNAIC{})


	//Rule tables
	DBConn.AutoMigrate(&ruletables.DiscountRules{})



	if err != nil {
		panic("Automigrate to database Failed")
	}
	fmt.Println("Database Migrated")
}
