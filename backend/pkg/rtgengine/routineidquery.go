package rtgengine

import (
	"github.com/harishm11/quoteCompare/database"
)

func GetRoutinId(rb string) string {
	var rout_id string
	db := database.DBConn
	db.Table("rate_routines").Where("ratebook_code = ? ", rb).Select("routine_id").Scan(&rout_id)
	return rout_id
}
