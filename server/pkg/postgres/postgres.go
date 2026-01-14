package postgres

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"Fytini/db"
)

func config() string {
	return "host=localhost user=gorm password=gorm dbname=gorm port=1234 sslmode=disable TimeZone=Asia/Bangkok"
}

func Open() (*gorm.DB, error) {
	dsn := config()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		// TODO: add logging and panic
		return nil, err
	}
	return db, nil
}

func Migrate(database *gorm.DB) error {
	if err := database.AutoMigrate(&db.Url{}); err != nil {
		// TODO: add logging and panic
		return err
	}
	return nil
}
