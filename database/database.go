package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	DBConn *gorm.DB
)

// import (
// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
//   )

//   dsn := "host=localhost user=postgres password=readonly dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
//   db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
