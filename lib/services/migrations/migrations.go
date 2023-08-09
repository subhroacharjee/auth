package migrations

import (
	"github.com/subhroacharjee/auth/lib/model/user"
	"gorm.io/gorm"
)

func MigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&user.User{})
}
