package handler

import (
	"net/http"
	"log"
	"encoding/json"
	"reminder_service/internal/reminder"
	"reminder_service/internal/service"

)

type Handler struct {
	s *service.Service;
}

func New(s *service.Service) *Handler {
	return &Handler{s: s};
}

func (h *Handler) handleAddReminder(w http.ResponseWriter, req *http.Request) {
	var r reminder.Reminder
	defer req.Body.Close()

	err := json.NewDecoder(req.Body).Decode(&r)
	if err != nil {
		log.Printf("Failed to decode reminder: %v", err)
		http.Error(w, "Invalid body.", http.StatusBadRequest)
		return
	}

	log.Printf("%v decoded reminder: %v", req.Method, r)

	h.s.AddReminder(r)

	log.Printf("cur reminders: %v", h.s.GetAllReminders());
}

func (h *Handler) handleListReminders(w http.ResponseWriter, req *http.Request) {
	rems := h.s.GetAllReminders()

	err := json.NewEncoder(w).Encode(rems);
	if err != nil {
		log.Printf("Failed to encode reminders: %v", err)
		http.Error(w, "Internal error.", http.StatusInternalServerError)
	}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /reminder", h.handleAddReminder)
	mux.HandleFunc("GET  /reminder", h.handleListReminders)
}
