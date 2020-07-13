package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
}