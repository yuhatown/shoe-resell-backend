package model

import (
	"fmt"
	"log"
	
	_ "github.com/joho/godotenv/autoload"
	"github.com/joho/godotenv"
)

func GetInfo() {
	db, err := ConnectDB()

	var str [6]string
	err = db.QueryRow("SELECT * FROM undefined").Scan(&str)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}