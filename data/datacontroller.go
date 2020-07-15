package data

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"nano/model"
)
var DB *gorm.DB

func InitDB()*gorm.DB{

	//dirvername :="mysql"
	//host :="localhost"
	//port :="3306"
	//database :="mycan"
	//uesrname :="root"
	//password :="root"
	//charset :="utf8"
	//args  :=fmt.Sprintf( "%s:%s@(%s:%s)/%s?charset=%s&parseTime=True",
	//	uesrname,
	//	password,
	//	host,
	//	port,
	//	database,
	//	charset,
	//)
	dsn := "root:root@tcp(127.0.0.1:3306)/mycan?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)


	if err !=nil{
		fmt.Println("fail to ")
		panic("fail to connect"+err.Error())
	}
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Todo {})
	db.AutoMigrate(&model.Webcount {})
	db.AutoMigrate(&model.Allweb{})
	db.AutoMigrate(&model.Webimage{})

	db.Create(&model.Webimage{Website:"http://www.airpano.com/?n=10&sort_by=&page=1" ,Image:"2016风景.png"})
	db.Create(&model.Webimage{Website:"http://www.airpano.com/?n=10&sort_by=&page=1" ,Image:"2018表情包"})

	DB=db
	return DB

}

func GetDB() *gorm.DB  {
	return DB
}