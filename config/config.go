package config

import (
	. "arvantest-monolith/DiscountSystem/Entity"
	. "arvantest-monolith/WalletSystem/Entity"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func GetDB() (*gorm.DB, error) {
	dbDriver := "sqlite3"
	dbName := "demo.db"

	db, err := gorm.Open(dbDriver, dbName)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&UserEntity{},&VoucherEntity{})
	return db, nil
}
