package rtgengine

import (
	"fmt"

	"github.com/harishm11/quoteCompare/database"
	"github.com/harishm11/quoteCompare/pkg/models"
	"github.com/harishm11/quoteCompare/pkg/tables"
)

func RatebookSelector(q *models.Quote) {
	db := database.DBConn
	var result tables.Ratebook
	db.Table("ratebooks").Where("ratebook_activation_date <= ? AND ratebook_new_business_eff_date <= ? ",
		q.RateAppliedDate, q.QuoteEffDate).Order("ratebook_activation_date desc").Last(&result)
	fmt.Println(result.RatebookCode)
	fmt.Println(q.RateAppliedDate)
	fmt.Println(q.QuoteEffDate)

}
