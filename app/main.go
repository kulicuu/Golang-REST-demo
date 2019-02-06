package main

import (
	"fmt"
	"net/http"
	"log"
)

// func handle(w http.ResponseWriter, r *http.Request) {
//
// 	fmt.Fprintf(w, "Tu have vieeesited %s in `%s`!", r.URL.Path, AppName)
//
// }

const AppName = "app"


func StartServer(port string) {
	// http.HandleFunc("/", handlerFunc)

	router := NewRouter()
	fmt.Println("Starting web server on port " + port)
	err := http.ListenAndServe(":" + port, router)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	StartServer("8080")
}
