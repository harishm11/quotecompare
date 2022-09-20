package rtgengine

import (
	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/tables/models"
	"github.com/harishm11/quoteCompare/tables/ratingtables"
)

func RatebookSelector(q *models.Quote) string {
	db := database.DBConn
	var result ratingtables.Ratebooks
	db.Table("ratebooks").Where("ratebook_activation_date <= ? AND ratebook_new_business_eff_date <= ? ",
		q.RateAppliedDate, q.QuoteEffDate).Order("ratebook_activation_date desc").Last(&result)
	return result.RatebookCode
}
