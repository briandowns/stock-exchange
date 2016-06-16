package main

import (
	"log"

	"github.com/briandowns/stock-exchange/database"
)

func main() {
	db, err := database.NewDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
