package main

import (
	"fmt"
	"github.com/MasaruFukazawa/go-de-smap-coding-challenge/database"
	"github.com/MasaruFukazawa/go-de-smap-coding-challenge/models"
)

func main() {
	fmt.Println("DBマイグレーション : Start")

	conn, err := database.NewDBConnection("./db.sqlite")

	defer conn.GetConnection()

	if err != nil {
		fmt.Println("DBマイグレーション : Error")
		return
	}

	err = conn.GetConnection().AutoMigrate(&models.User{}, &models.Consumption{})

	fmt.Println("DBマイグレーション : Success")
	return
}
