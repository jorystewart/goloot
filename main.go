package main

import (
  "database/sql"
  "fmt"
  _ "modernc.org/sqlite"
  //"net/http"
)

/*func computersHandler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Computers")
}*/

func main() {
  db, err := sql.Open("sqlite", "jdbc:sqlite:D:\\source\\repos\\jorystewart\\goloot\\data\\roster.sqlite")
  if err != nil {
    fmt.Sprintln("Unable to connect to database")
    return
  }
  defer db.Close()

  queryResult, _ := db.Query("SELECT * FROM roster WHERE name = 'Anthalon'")
  for queryResult.Next() {
    fmt.Println(queryResult.Scan())
  }

  //http.HandleFunc("/computers", computersHandler)
  //http.ListenAndServe(":5050", nil)
}
