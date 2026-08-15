package main

import (
	"net/http"
	"os"
	"log"
	"reminder_service/internal/handler"
	"reminder_service/internal/service"
	"reminder_service/internal/reminder"
)

func main() {
	addr, _ := os.LookupEnv("ADDR")

	repo := reminder.NewRepo()
	svc := service.New(repo)
	hndl := handler.New(svc)

	mux := http.NewServeMux()
	hndl.RegisterRoutes(mux)

	err := http.ListenAndServe(addr, mux)
	if err != nil {
		log.Println("Failed to create HTTP server: ", err)
	}
}
