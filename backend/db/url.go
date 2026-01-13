package db

import gorm.io/gorm

type Url struct {
	gorm.Model
	Url string
	expiration_date string
	custom_alias
}
