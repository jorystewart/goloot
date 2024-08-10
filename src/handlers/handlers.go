package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Test struct {
	Class string `json:"class" bson:"Class"`
}

func ClassesHandler(w http.ResponseWriter, r *http.Request) {
	var requestContainer Test
	fmt.Fprintln(w, "ClassesHandler")
	fmt.Fprintln(w, r.Method)
	fmt.Fprintln(w, r.Header)
	fmt.Fprintln(w, r.Body)
	fmt.Println("ClassesHandler")
	fmt.Println(r.Method)
	fmt.Println(r.Header)
	err := json.NewDecoder(r.Body).Decode(&requestContainer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(requestContainer.Class)
}
