package quote

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/harishm11/quoteCompare/database"
	"gorm.io/gorm"
)

type Quote struct {
	gorm.Model
	QuoteformFields []struct {
		gorm.Model
		QuoteNumber int    `json:"quoteNumber"`
		Lob         string `json:"lob"`
	} `json:"quoteformFields"`
	DriverformFields []struct {
		gorm.Model
		Name         string `json:"name"`
		Age          string `json:"age"`
		Experience   string `json:"experience"`
		Course       string `json:"course"`
		Incidentdate string `json:"incidentdate"`
		Incidenttype string `json:"incidenttype"`
	} `json:"driverformFields"`
	VehicleformFields []struct {
		gorm.Model
		Year          string `json:"year"`
		Make          string `json:"make"`
		VModel        string `json:"model"`
		AnnualMileage string `json:"annualMileage"`
		GrgZip        string `json:"grgZip"`
	} `json:"vehicleformFields"`
}

func GetQuotes(c *fiber.Ctx) error {

	db := database.DBConn
	var quotes []Quote
	db.Find(&quotes)
	return c.JSON(quotes)
}

func GetQuote(c *fiber.Ctx) error {
	id := c.Params("QuoteNumber")
	db := database.DBConn
	var quote Quote
	db.Find(&quote, id)
	return c.JSON(quote)
}

func NewQuote(c *fiber.Ctx) error {
	db := database.DBConn

	quote := new(Quote)

	if err := c.BodyParser(quote); err != nil {
		return err
	}
	fmt.Println(quote.QuoteformFields)
	fmt.Println(quote.DriverformFields)
	// // VehicleData := make(map[veh]vehicle{})

	//DTO for rating variables - parent process - P2 - create generic and company specific rating variables
	//Concurrent 100 company quotes - child processes - PD
	//each child process for compaany gets calc rule for each coverage from redis cache , get the factors , calc the premium, send premium response - P3

	db.Create(&quote)
	return c.JSON(quote)
}

func DeleteQuote(c *fiber.Ctx) error {
	id := c.Params("QuoteNumber")
	db := database.DBConn

	var quote Quote
	db.First(&quote, id)

	db.Delete(&quote)
	return c.SendString("Deleted")
}
