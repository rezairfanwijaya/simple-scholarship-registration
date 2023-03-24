package main

import (
	"log"
	"serkom/database"
)

func main() {
	conn, err := database.NewConnection(".env")
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Println(conn)
}
