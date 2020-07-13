package data

import (

	"fmt"
	"github.com/jinzhu/gorm"
	_"github.com/go-sql-driver/mysql"

	"nano/model"
)
var DB *gorm.DB
func InitDB()*gorm.DB{

	dirvername :="mysql"
	host :="localhost"
	port :="3306"
	database :="demo"
	uesrname :="root"
	password :="root"
	charset :="utf8"
	args  :=fmt.Sprintf( "%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
		uesrname,
		password,
		host,
		port,
		database,
		charset,
	)
	db,err :=gorm.Open(dirvername,args)

	if err !=nil{
		panic("fail to connect"+err.Error())
	}
	db.AutoMigrate(&model.User{})

	//log.Printf("2333")

	DB=db
	return DB

}

func GetDB() *gorm.DB  {
	return DB
}