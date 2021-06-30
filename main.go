package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main(){
	log.Println("Helloworld")
	dsn := "food_delivery:19e5a718a54a9fe0559dfbce6908@tcp(127.0.0.1:3306)/food_delivery?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		log.Fatalln("[err]MySql:", err)
	}
	log.Println(db)
}