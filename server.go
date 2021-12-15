package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World")
	})

	fmt.Println("open http://localhost:9999")
	http.ListenAndServe(":9999", nil)
}
