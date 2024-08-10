package data

import (
	"database/sql"
	"fmt"
	_ "modernc.org/sqlite"
	"strings"
)

type RosterMember struct {
	name   string
	class  string
	isMain bool
}

func QueryRosterClass(db *sql.DB, class string) ([]RosterMember, error) {
	var results []RosterMember
	dbStatus := db.Ping()
	if dbStatus != nil {
		return results, dbStatus
	}

	queryResult, err := db.Query("SELECT * FROM roster WHERE class = ?", class)
	if err != nil {
		return results, err
	}

	if queryResult == nil {
		return results, nil
	} else {
		for queryResult.Next() {
			member := RosterMember{}
			err := queryResult.Scan(&member.name, &member.class, &member.isMain)
			if err != nil {
				return results, err
			}
			results = append(results, member)
		}
		return results, nil
	}
}

func QueryRosterName(db *sql.DB, names []string) ([]RosterMember, error) {
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
			err := queryResult.Scan(&member.name, &member.class, &member.isMain)
			if err != nil {
				fmt.Println(err)
			}
			results = append(results, member)
		}
		return results, nil
	}
}
