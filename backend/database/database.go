package database

import (
	"fmt"

	"github.com/harishm11/quoteCompare/pkg/models"
	"github.com/harishm11/quoteCompare/pkg/tables"
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

	DBConn.AutoMigrate(&models.Quote{})
	DBConn.AutoMigrate(&models.Driver{})
	DBConn.AutoMigrate(&models.Vehicle{})
	DBConn.AutoMigrate(&tables.Ratebooks{})
	DBConn.AutoMigrate(&tables.Routines{})
	DBConn.AutoMigrate(&tables.RateFactors{})
	DBConn.AutoMigrate(&tables.RateRoutinSteps{})

	if err != nil {
		panic("Automigrate to database Failed")
	}
	fmt.Println("Database Migrated")
}
