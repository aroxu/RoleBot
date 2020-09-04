package initializer

import (
	"B1ackAnge1/RoleBot/model"
	"B1ackAnge1/RoleBot/utils"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() error {
	db, errFailedOpenDatabase := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if errFailedOpenDatabase != nil {
		log.Fatalln("Failed to open database.")
		return errFailedOpenDatabase
	}
	utils.SetDB(db)
	log.Print("Successfully Connected To Database")

	var models = []interface{}{&model.Vote{}}
	errFailedAutoMigrate := db.AutoMigrate(models...)
	if errFailedAutoMigrate != nil {
		log.Fatalln("Failed to perform AutoMigrate.")
		return errFailedAutoMigrate
	}
	log.Print("Successfully performed AutoMigrate")
	return nil
}
