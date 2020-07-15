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

	db.Create(&model.Webimage{Website:"http://www.airpano.com" ,Image:"2016风景.png"})
	db.Create(&model.Webimage{Website:"https://sorry.xuty.tk" ,Image:"2018表情包.png"})
	db.Create(&model.Webimage{Website:"http://www.ilidilid.com" ,Image:"2018反向.png"})
	db.Create(&model.Webimage{Website:"https://photomosh.com" ,Image:"2018在线效果.png"})
	db.Create(&model.Webimage{Website:"https://thetruesize.com" ,Image:"2018真实大小.png"})
	db.Create(&model.Webimage{Website:"https://zh.wix.com" ,Image:"2018做网站.png"})
	db.Create(&model.Webimage{Website:"https://resn.co.nz" ,Image:"2019？.png"})
	db.Create(&model.Webimage{Website:"https://creativemass.cn" ,Image:"2019创意.jpg"})
	db.Create(&model.Webimage{Website:"https://www.ls.graphics" ,Image:"2019渐变图.png"})
	db.Create(&model.Webimage{Website:"https://cavalierchallenge.com" ,Image:"2019马.png"})
	db.Create(&model.Webimage{Website:"https://sf.taobao.com/?spm=a213w.7398504.sfhead2014.2.9vuB5l&current=index" ,Image:"2019拍卖.png"})
	db.Create(&model.Webimage{Website:"https://app.grammarly.com" ,Image:"2019grammerly.png"})
	db.Create(&model.Webimage{Website:"http://listen1.github.io/listen1" ,Image:"2019listen1.png"})
	db.Create(&model.Webimage{Website:"https://leptc.github.io/bili" ,Image:"2019up主排名.png"})
	db.Create(&model.Webimage{Website:"https://mj.yuzhua.com/search/3.html" ,Image:"2019yuzhua.png"})
	db.Create(&model.Webimage{Website:"https://colourise.sg" ,Image:"2020老照片.png"})
	db.Create(&model.Webimage{Website:"https://bongo.cat" ,Image:"2020猫.png"})
	DB=db
	return DB

}

func GetDB() *gorm.DB  {
	return DB
}