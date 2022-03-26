package sqlite

import (
	"spa-go/internal/admin"
	"spa-go/internal/user"
	"spa-go/internal/utils/db"
)

func AotoMigrate() {
	db.Sqlite.AutoMigrate(&admin.Admin{})
	db.Sqlite.AutoMigrate(&user.User{})
}
