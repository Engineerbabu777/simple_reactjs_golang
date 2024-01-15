package main

import (
	"encoding/json"
	"net/http"
)

type ApiResponse struct {
    Message string `json:"message"`
}

func main() {
    http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
        res.Write([]byte("Hello from Golang server!"))
    })

	http.HandleFunc("/api/data", dataHandler)


    http.ListenAndServe(":8080", nil)
}


func dataHandler(res http.ResponseWriter, req *http.Request){
	response := ApiResponse{Message: "Data from the API"};
	json.NewEncoder(res).Encode(response);
}