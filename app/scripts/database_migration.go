package main

import (
	"fmt"
	"os"
	"github.com/joho/godotenv"
	"github.com/MasaruFukazawa/go-de-smap-coding-challenge/database"
	"github.com/MasaruFukazawa/go-de-smap-coding-challenge/domain/entities"
)

func main() {
	fmt.Println("DBマイグレーション : Start")

	err := godotenv.Load()

	if err != nil {
		fmt.Println("DBマイグレーション : Error loading .env file")
		return
	}

	db, err := database.NewDB(os.Getenv("SQLITE3_FILE_PATH"))

	if err != nil {
		fmt.Println("DBマイグレーション : Error")
		return
	}

	defer db.Close()

	err = db.GetConnection().AutoMigrate(&entities.User{}, &entities.Consumption{})

 	if err != nil {
    	fmt.Println("DBマイグレーション : Error -", err)
     	return
  	}

	fmt.Println("DBマイグレーション : Success")

	return
}
