package main

import (
	"log"
	"net/http"

	"github.com/vivek6201/go-api-gateway/internals/router"
)

func main() {
	r := router.NewRouter()

	log.Println("gateway server is running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
