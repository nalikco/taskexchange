package service

import (
	"taskexchange"
	"taskexchange/pkg/repository"
	"time"
)

type EventsService struct {
	repo repository.Events
}

func NewEventsService(repo repository.Events) *EventsService {
	return &EventsService{
		repo: repo,
	}
}

func (s *EventsService) CreateEvent(userId int, message, link string) (int, error) {
	event := taskexchange.Event{
		UserId:    userId,
		Message:   message,
		Link:      link,
		CreatedAt: time.Now(),
	}

	id, err := s.repo.Create(event)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (s *EventsService) PollingEvents(userId, id int) ([]taskexchange.Event, error) {
	events, err := s.repo.FindPolling(userId, id)
	if err != nil {
		return []taskexchange.Event{}, err
	}

	return events, nil
}

func (s *EventsService) GetNewEvents(userId int) ([]taskexchange.Event, error) {
	return s.repo.FindNew(userId)
}

func (s *EventsService) GetLastUserEventId(userId int) (int, error) {
	var id int
	lastUserEvent, err := s.repo.FindLastUser(userId)
	if err != nil && err.Error() != "sql: no rows in result set" {
		return 0, err
	}

	id = lastUserEvent.ID

	return id, nil
}

func (s *EventsService) ViewEvent(userId, id int) error {
	return s.repo.View(userId, id)
}

func (s *EventsService) DeleteEvent(userId, id int) error {
	return s.repo.Delete(userId, id)
}
