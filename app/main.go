package main

import (
	"app/routes"
	"fmt"
	"log"
)

func main() {
	r := routes.SetupRouter()

	fmt.Println("Server is running on port 8080...")
	log.Fatal(r.Run(":8080"))
}
