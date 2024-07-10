package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Define a struct to hold the POST request data
type Data struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	http.HandleFunc("/", getHandler)
	http.HandleFunc("/data", postHandler)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

// GET request handler
func getHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "So wake me up when it's all over When I'm wiser and I'm older")
}

// POST request handler
func postHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	var data Data
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	response := fmt.Sprintf("Received data: Name = %s, Age = %d", data.Name, data.Age)
	fmt.Fprintf(w, response)
}
