package quote

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/harishm11/quoteCompare/database"
	"github.com/jinzhu/gorm"
)

type Quote struct {
	gorm.Model
	QuoteNumber int    `json:"quoteNumber"`
	Lob         string `json:"lob"`
	Drivers     struct {
		gorm.Model
		Driver []struct {
			gorm.Model
			DrvName             string `json:"drvName"`
			DrvDrivingEperience int    `json:"drvDrivingEperience"`
			DrvAgeInYrs         int    `json:"drvAgeInYrs"`
			DrvDrivingCourse    bool   `json:"drvDrivingCourse"`
		} `json:"driver"`
	} `json:"drivers"`
	Vehicles struct {
		gorm.Model
		Vehicle []struct {
			gorm.Model
			VehYear          int    `json:"vehYear"`
			VehMake          string `json:"vehMake"`
			VehModel         string `json:"vehModel"`
			VehAnnualMileage int    `json:"vehAnnualMileage"`
			VehGrgZipcode    string `json:"vehGrgZipcode"`
		} `json:"vehicle"`
	} `json:"vehicles"`
	Incidents struct {
		gorm.Model
		Incident []struct {
			gorm.Model
			IncDate   string `json:"incDate"`
			IncType   string `json:"incType"`
			Incdriver string `json:"incdriver"`
		} `json:"incident"`
	} `json:"incidents"`
}

// type Quote struct {
// 	gorm.Model
// 	QuoteNumber int       `json:"quoteNumber"`
// 	Lob         string    `json:"lob"`
// 	Drivers     Drivers   `json:"drivers"`
// 	Vehicles    Vehicles  `json:"vehicles"`
// 	Incidents   Incidents `json:"incidents"`
// }
// type Driver struct {
// 	gorm.Model
// 	DrvName             string `json:"drvName"`
// 	DrvDrivingEperience int    `json:"drvDrivingEperience"`
// 	DrvAgeInYrs         int    `json:"drvAgeInYrs"`
// 	DrvDrivingCourse    bool   `json:"drvDrivingCourse"`
// }

// type Drivers struct {
// 	Driver []Driver `json:"driver"`
// }
// type Vehicle struct {
// 	gorm.Model
// 	VehYear          int    `json:"vehYear"`
// 	VehMake          string `json:"vehMake"`
// 	VehModel         string `json:"vehModel"`
// 	VehAnnualMileage int    `json:"vehAnnualMileage"`
// 	VehGrgZipcode    string `json:"vehGrgZipcode"`
// }

// type Vehicles struct {
// 	Vehicle []Vehicle `json:"vehicle"`
// }
// type Incident struct {
// 	gorm.Model
// 	IncDate   string `json:"incDate"`
// 	IncType   string `json:"incType"`
// 	Incdriver string `json:"incdriver"`
// }

// type Incidents struct {
// 	Incident []Incident `json:"incident"`
// }

func GetQuotes(c *fiber.Ctx) {
	db := database.DBConn
	var quotes []Quote
	db.Find(&quotes)
	c.JSON(quotes)
}

func GetQuote(c *fiber.Ctx) {
	id := c.Params("QuoteNumber")
	db := database.DBConn
	var quote Quote
	db.Find(&quote, id)
	c.JSON(quote)
}

func NewQuote(c *fiber.Ctx) {
	db := database.DBConn

	quote := new(Quote)

	if err := c.BodyParser(quote); err != nil {
		c.Status(503).Send(err)
		return
	}
	fmt.Println(quote.QuoteNumber)
	fmt.Println(quote.Vehicles.Vehicle[1].VehYear)
	VehicleData := make(map[veh]vehicle{})

	// PAratingVariables.drvvar()

	//DTO for rating variables - parent process - P2 - create generic and company specific rating variables
	//Concurrent 100 company quotes - child processes - PD
	//each child process for compaany gets calc rule for each coverage from redis cache , get the factors , calc the premium, send premium response - P3

	db.Create(&quote)
	c.JSON(quote)
}

func DeleteQuote(c *fiber.Ctx) {
	id := c.Params("QuoteNumber")
	db := database.DBConn

	var quote Quote
	db.First(&quote, id)
	if quote.Lob == "" {
		c.Status(500).Send("No quote found with given QuoteNumber")
		return
	}
	db.Delete(&quote)
	c.Send("Quote successfully deleted")
}
