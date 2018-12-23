package main

import (
	"fmt"
	"log"
	"net/http"

	// "time"
	"moff-backend/databases"
	"moff-backend/routes"
)

func test(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}

func main() {
	databases.Init()
	http.HandleFunc("/test", test)
	http.Handle("/", routes.AllRoutes())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
