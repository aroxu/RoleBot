package utils

import "gorm.io/gorm"

var (
	gDb     *gorm.DB
)

func SetDB(db *gorm.DB) {
	gDb = db
}

func GetDB() *gorm.DB {
	return gDb
}
