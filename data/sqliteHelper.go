package data

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"strings"
)

type RosterMember struct {
	Name   string `json:"name"`
	Class  string `json:"class"`
	IsMain bool   `json:"isMain"`
}

var db *sql.DB

func InitializeDbConnection() error {
	var err error
	db, err = sql.Open("sqlite", "data/roster.sqlite")
	if err != nil {
		return err
	}
	return db.Ping()
}

func QueryRosterClass(class string) ([]RosterMember, error) {
	var results []RosterMember
	dbStatus := db.Ping()
	if dbStatus != nil {
		return nil, dbStatus
	}

	queryResult, err := db.Query("SELECT * FROM roster WHERE upper(class) = upper(?)", class)
	if err != nil {
		return results, err
	}

	if queryResult == nil {
		return results, nil
	} else {
		for queryResult.Next() {
			member := RosterMember{}
			err := queryResult.Scan(&member.Name, &member.Class, &member.IsMain)
			if err != nil {
				return results, err
			}
			results = append(results, member)
		}
		return results, nil
	}
}

func QueryRosterName(names []string) ([]RosterMember, error) {
	var results []RosterMember
	dbStatus := db.Ping()
	if dbStatus != nil {
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
	}

	queryResult, err := db.Query(stringBuilder.String(), test...)
	if err != nil {
		return results, err
	}

	if queryResult == nil {
		return results, nil
	} else {
		for queryResult.Next() {
			member := RosterMember{}
			err := queryResult.Scan(&member.Name, &member.Class, &member.IsMain)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, member)
		}
		return results, nil
	}
}
