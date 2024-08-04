package src

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"net/http"
	"strings"
)

func classesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ClassesHandler")
}

type RosterMember struct {
	name   string
	class  string
	isMain bool
}

func main() {

	db, err := sql.Open("sqlite", "../data/roster.sqlite")
	if err != nil {
		fmt.Sprintln("Unable to connect to database")
		return
	}
	defer db.Close()

	classTest, err := queryRosterClass(db, "Warrior")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Class query test:")
		fmt.Println(classTest)
	}

	names := []string{"Anthalon", "Ynystere"}
	nameTest, err := queryRosterName(db, names)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Name query test:")
		fmt.Println(nameTest)
	}

	http.HandleFunc("/classes", classesHandler)
	http.ListenAndServe(":5050", nil)
}

func queryRosterClass(db *sql.DB, class string) ([]RosterMember, error) {
	var results []RosterMember
	dbStatus := db.Ping()
	if dbStatus != nil {
		fmt.Println("Unable to connect to database")
		return results, dbStatus
	}

	queryResult, err := db.Query("SELECT * FROM roster WHERE class = ?", class)
	if err != nil {
		fmt.Sprintln("Unable to query database")
		return results, err
	}

	if queryResult == nil {
		fmt.Println("No results found")
		return results, nil
	} else {
		for queryResult.Next() {
			member := RosterMember{}
			err := queryResult.Scan(&member.name, &member.class, &member.isMain)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, member)
		}
		return results, nil
	}
}

func queryRosterName(db *sql.DB, names []string) ([]RosterMember, error) {
	var results []RosterMember
	dbStatus := db.Ping()
	if dbStatus != nil {
		fmt.Println("Unable to connect to database")
		return results, dbStatus
	}

	stringBuilder := strings.Builder{}
	stringBuilder.WriteString("SELECT * FROM roster WHERE name IN (?")
	for i := 0; i < len(names)-1; i++ {
		stringBuilder.WriteString(", ?")
	}
	stringBuilder.WriteString(")")

	test := make([]interface{}, len(names))
	for i, v := range names {
		test[i] = v
		fmt.Println(v)
	}

	queryResult, err := db.Query(stringBuilder.String(), test...)
	if err != nil {
		fmt.Sprintln("Unable to query database")
		return results, err
	}

	if queryResult == nil {
		fmt.Println("No results found")
		return results, nil
	} else {
		for queryResult.Next() {
			member := RosterMember{}
			err := queryResult.Scan(&member.name, &member.class, &member.isMain)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, member)
		}
		return results, nil
	}
}
