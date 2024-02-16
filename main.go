package main

import (
	"net/http"

	"github.com/gregoryAlvim/api-rest-in-go/routes"
)

func main() {
	router := routes.New()

	http.ListenAndServe(":3000", router)
}
