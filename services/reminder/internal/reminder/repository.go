// Package reminder implements ReminderRepository storage and Reminder struct.
package reminder

import (
	"sync"
)

func NewRepo() *ReminderRepo {
	return &ReminderRepo{};
}

type ReminderRepo struct {
	rems []Reminder;
	mu sync.Mutex;
}

func (repo *ReminderRepo) Add(r Reminder) {
	repo.mu.Lock();
	defer repo.mu.Unlock()

	repo.rems = append(repo.rems, r)

}

func (repo *ReminderRepo) GetAll() []Reminder {
	repo.mu.Lock()
	defer repo.mu.Unlock()

	ret := make([]Reminder, len(repo.rems))
	copy(ret, repo.rems)

	return ret
}
