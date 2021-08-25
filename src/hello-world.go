package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/", HelloServer)
	fmt.Printf("http://0.0.0.0:8080\n")
	http.ListenAndServe(":8080", nil)
}

func HelloServer(w http.ResponseWriter, r *http.Request) {
	name, err := os.Hostname()
	if err != nil {
		name = "get hostname failed"
	}
	fmt.Fprintf(w, "Hello, %s!", name)
}
