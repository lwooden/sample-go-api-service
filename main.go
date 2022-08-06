package main

import (
	"sample-go-api-service/routes"
)

func main() {

	// Setup routes
	r := routes.SetupRouter()

	// Running
	r.Run(":8080")

}
