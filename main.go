package main

import (
	"fmt"
	"net/http"
	"os"
)

func BaseHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func main() {
	addr, _ := os.LookupEnv("ADDR")

	http.HandleFunc("/", BaseHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		fmt.Println("Failed to create HTTP server: ", err)
	}
}
