// Package service implements reminder service business logic.
package service

import (
	"reminder_service/internal/reminder"
)

type Service struct {
	repo *reminder.ReminderRepo;
}

func New(repo *reminder.ReminderRepo) *Service {
	return &Service{repo: repo}
}

func (s *Service) AddReminder(r reminder.Reminder) {
	s.repo.Add(r)
}

func (s *Service) GetAllReminders() []reminder.Reminder {
	return s.repo.GetAll()
}
