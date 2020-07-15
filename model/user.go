package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;unique"`
	Password string `gorm:"type:varchar(100);not null"`
}
type Todo struct {
	ID int `json:"id"`
	Title string `json:"title"`
	ST bool `json:"st"`
}
type Webcount struct {
	gorm.Model
	Name string `gorm:"type:varchar(20);not null;"`
	Website string `gorm:"type:varchar(100);not null"`
	Introduction string `gorm:"type:varchar(200);not null"`

}
type Allweb struct {
	gorm.Model

	Website string `gorm:"type:varchar(100);not null;unique"`
	Num int     `gorm:"AUTO_INCREMENT"`
}
type Webimage struct {
	gorm.Model

	Website string `gorm:"type:varchar(100);not null;unique"`
	Image string     `gorm:"type:varchar(100);not null"`
}
