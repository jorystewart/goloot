package main

import (
	"fmt"
	"goloot/data"
	"goloot/src/handlers"
	_ "modernc.org/sqlite"
	"net/http"
)

func main() {
	err := data.InitializeDbConnection()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Connected to database")
	}

	http.HandleFunc("/class", handlers.ClassHandler)
	http.HandleFunc("/name", handlers.NameHandler)
	http.ListenAndServe(":5050", nil)
}
