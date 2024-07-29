package main

import (
  "database/sql"
  "fmt"
  _ "modernc.org/sqlite"
  "net/http"
)

func computersHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Computers")
}

func main() {
  http.HandleFunc("/computers", computersHandler)
  http.ListenAndServe(":5050", nil)
}

func connect_db(name string) *sql.DB {
  db, err := sql.Open(name, "jdbc:sqlite:D:\\source\\repos\\jorystewart\\goloot\\data\\roster.sqlite")
  if err != nil {
    fmt.Sprintln("Unable to connect to database")
  }
  return db
}
