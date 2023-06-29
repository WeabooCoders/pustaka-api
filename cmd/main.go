package main

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	dsn := "host=172.17.0.2 user=root password=root dbname=perpus port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err.Error())
	}

	fmt.Println("oke")

}
