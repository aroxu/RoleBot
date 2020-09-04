package utils

import "gorm.io/gorm"

var (
	gDb     *gorm.DB
)

//SetDB sets database config from ceasy.config.json
func SetDB(db *gorm.DB) {
	gDb = db
}

//GetDB gets database config from ceasy.config.json
func GetDB() *gorm.DB {
	return gDb
}
