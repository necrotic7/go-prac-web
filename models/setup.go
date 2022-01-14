package models

import(
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const(
	USERNAME = "root"
	PASSWORD = "80806789"
	NETWORK = "tcp"
	SERVER = "127.0.0.1"
	PORT = 3306
	DATABASE = "demo"
)

var DB *gorm.DB

func ConnectDataBase(){
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",USERNAME,PASSWORD,NETWORK,SERVER,PORT,DATABASE)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("connection error:"+err.Error())
	}else{
		fmt.Println("Connect success")
	}

	db.AutoMigrate(&Book{})
	db.AutoMigrate(&User{})
	DB = db
}