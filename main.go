package main

import (
	"fmt"
	"github.com/4damrr/template_test/template"
	"net/http"
)

func main() {
	// Route setup
	http.HandleFunc("/", template.GenerateCV)

	// Start the server on port 8080
	fmt.Println("Server running at http://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Server failed:", err)
	}
}
