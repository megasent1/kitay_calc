package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	html, err := os.ReadFile("index.html")
	if err != nil {
		fmt.Println("Error reading index.html:", err)
		return
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, string(html))
	})

	fmt.Println("Server starting on port 8457")
	http.ListenAndServe(":8457", nil)
}
