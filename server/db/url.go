package db

import (
	"time"

	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url            string
	ExpirationDate *time.Time
	CustomAlias    string
}
