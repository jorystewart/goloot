package handlers

import (
	"encoding/json"
	"fmt"
	"goloot/data"
	"net/http"
)

type QueryContainer struct {
	Name  []string `json:"name" bson:"name"`
	Class string   `json:"class" bson:"class"`
}

func ClassHandler(resp http.ResponseWriter, req *http.Request) {
	var requestContainer QueryContainer
	err := json.NewDecoder(req.Body).Decode(&requestContainer)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
		resp.Header().Set("Content-Type", "application/json")
		errorBody := []byte(`{"error":"query failed"}`)
		resp.Write(errorBody)
	}

	if req.Method == "GET" {
		queryResult, err := data.QueryRosterClass(requestContainer.Class)
		if err != nil {
			fmt.Println("Query failed")
			fmt.Println(err)
			resp.WriteHeader(http.StatusBadRequest)
			resp.Header().Set("Content-Type", "application/json")
			errorBody := []byte(`{"error":"query failed"}`)
			resp.Write(errorBody)
		} else {
			jsonResponse, _ := json.Marshal(queryResult)
			resp.WriteHeader(http.StatusOK)
			resp.Header().Set("Content-Type", "application/json")
			resp.Write(jsonResponse)
		}
	}
}

func NameHandler(resp http.ResponseWriter, req *http.Request) {
	var requestContainer QueryContainer
	err := json.NewDecoder(req.Body).Decode(&requestContainer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(requestContainer.Name)

	if req.Method == "GET" {
		fmt.Println("GET")
		queryResult, err := data.QueryRosterName(requestContainer.Name)
		if err != nil {
			fmt.Println("Query failed")
			fmt.Println(err)
		} else {
			fmt.Println("Query succeeded")
			fmt.Println(queryResult)
		}
	}
}
