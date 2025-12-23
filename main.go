package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("First Program")
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "STATUS: OK")
	})

	http.ListenAndServe(":8080", nil)
}
