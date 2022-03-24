package sqlite

import (
	"spa-go/internal/utils/db"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func Init() {
	var err error
	db.Sqlite, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic("failed to connect database")
	}
	AotoMigrate()
}
