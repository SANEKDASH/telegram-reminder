package main

import (
	"fmt"
	"net/http"
	"os"
	"log"
	"sync"
	"encoding/json"
)

type Reminder struct {
	Msg string `json:"msg"`
	Time string `json:"time"`
}

type Reminders struct {
	rms []Reminder;
	mu sync.Mutex;
}

var reminders Reminders;

func BaseHandler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Hello!")
}

func ListReminders(w http.ResponseWriter, req *http.Request) {
	reminders.mu.Lock()
	defer reminders.mu.Unlock()

	err := json.NewEncoder(w).Encode(reminders.rms);
	if err != nil {
		http.Error(w, "Internal error.", http.StatusInternalServerError)
	}
}

func AddReminder(w http.ResponseWriter, req *http.Request) {
	var r Reminder;
	defer req.Body.Close()

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		http.Error(w, "Invalid body.", http.StatusBadRequest)
		return
	}

	log.Printf("%v reminder: %v", req.Method, r)

	reminders.mu.Lock()
	defer reminders.mu.Unlock()
	reminders.rms = append(reminders.rms, r);


	log.Printf("cur reminders: %v", reminders.rms);
}

func ReminderHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		ListReminders(w, req)
	case http.MethodPost:
		AddReminder(w, req)
		fmt.Fprintf(w, "Post method.")
	default:
		log.Printf("Bad request: %v", req)
		http.Error(w, "Invalid method.", http.StatusBadRequest)
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
