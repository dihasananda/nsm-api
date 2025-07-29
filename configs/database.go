package configs

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	_ "modernc.org/sqlite" // import pure Go sqlite
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Dialector{
		DriverName: "sqlite", // uses modernc.org/sqlite under the hood
		DSN:        "./db/nsms.db",
	}, &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}
	return db
}
