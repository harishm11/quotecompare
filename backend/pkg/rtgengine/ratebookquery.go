package rtgengine

import (
	"time"

	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/tables/ratingtables"
)

func RatebookSelector(qeffdt time.Time, qapplieddt time.Time) (string, time.Time) {
	db := database.DBConn
	var result ratingtables.Ratebooks
	db.Table("ratebooks").Where("ratebook_activation_date <= ? AND ratebook_new_business_eff_date <= ? ",
		qapplieddt, qeffdt).Order("ratebook_activation_date desc").Last(&result)
	return result.RatebookCode, result.RatebookActivationDate
}
