package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
)

type Reminder struct {
	Msg string `json:"msg"`
	Time string `json:"time"`
}

func BaseHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello!")
}

func ReminderHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		fmt.Fprintf(w, "Get method.")
	case http.MethodPost:
		fmt.Fprintf(w, "Post method.")
	default:
		log.Printf("Bad request: %v", req)
		fmt.Fprintf(w, "Error %d", http.StatusBadRequest)
	}
}

func main() {
	addr, _ := os.LookupEnv("ADDR")

	http.HandleFunc("/", BaseHandler)
	http.HandleFunc("/reminder", ReminderHandler)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Println("Failed to create HTTP server: ", err)
	}
}
