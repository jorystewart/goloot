package main

import (
	"database/sql"
	"fmt"
	"goloot/data"
	"goloot/src/handlers"
	_ "modernc.org/sqlite"
	"net/http"
)

func main() {
	db, err := sql.Open("sqlite", "data/roster.sqlite")
	if err != nil {
		return
	}

	err = db.Ping()
	if err != nil {
		return
	}
	defer db.Close()

	classTest, err := data.QueryRosterClass(db, "Warrior")
	if err != nil {
		fmt.Print("queryRosterClass: ")
		fmt.Println(err)
	} else {
		fmt.Println("Class query test:")
		fmt.Println(classTest)
	}

	names := []string{"Anthalon", "Ynystere"}
	nameTest, err := data.QueryRosterName(db, names)
	if err != nil {
		fmt.Print("queryRosterName: ")
		fmt.Println(err)
	} else {
		fmt.Println("Name query test:")
		fmt.Println(nameTest)
	}

	http.HandleFunc("/classes", handlers.ClassesHandler)
	http.ListenAndServe(":5050", nil)
}
