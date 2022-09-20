package rtgengine

import (
	"fmt"

	"github.com/harishm11/quoteCompare/pkg/models"
)

func RatingEngineImpl(q *models.Quote) {

	//DTO for rating variables - parent process - P2 - create generic and company specific rating variables
	//Concurrent 100 company quotes - child processes - PD
	//each child process for compaany gets calc rule for each coverage from redis cache , get the factors , calc the premium, send premium response - P3

	var ratebookcode = RatebookSelector(q)
	fmt.Println(" Quote eff date ", q.QuoteEffDate)
	fmt.Println("Ratebook code", ratebookcode)
}
