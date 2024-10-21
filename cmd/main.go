package main

import (
	"log"
	
	"github.com/WilderEdwards/ecom-api-go/cmd/api"
)

func main() {
	//create new api endpoint
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
