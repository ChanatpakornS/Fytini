package db

import (
	"gorm.io/gorm"
)

type Url struct {
	gorm.Model
	Url            string
	ExpirationDate string
	CustomAlias    string
}
