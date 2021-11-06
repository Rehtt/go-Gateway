package models

import "github.com/jinzhu/gorm"

type UserTables struct {
	*gorm.Model
	Name string `gorm:"not null"`
}
