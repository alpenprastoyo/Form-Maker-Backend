package main

import (
	"form-api/routers"
	"os"
)

func main() {

	r := routers.SetupRouter()

	port := os.Getenv("port")
	if port == "" {
		port = "8050" //localhost
	}

	r.Run(":" + port)

}
