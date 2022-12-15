package event

import (
	"dev11/api/domain"
	"errors"
	"log"
	"sync"
	"time"
)

const (
	Nil = StoreError("cache: nil")
)

type StoreError string

func (e StoreError) Error() string { return string(e) }

type Store struct {
	mu     sync.RWMutex
	LastID int64
	Data   map[int64]domain.Event
	logger *log.Logger
}

func NewInmemStore(logger *log.Logger) *Store {
	var ims = Store{
		mu:     sync.RWMutex{},
		LastID: -1,
		Data:   map[int64]domain.Event{},
	}

	return &ims
}

func (s *Store) Put(event *domain.Event) int64 {

	s.mu.Lock()
	defer s.mu.Unlock()

	if event.ID > -1 {
		event.UpdatedAt = time.Now()
		s.Data[event.ID] = *event
	} else {
		s.LastID++
		event.CreatedAt = time.Now()
		event.ID = s.LastID
		s.Data[s.LastID] = *event
	}

	return s.LastID
}

func (s *Store) GetById(id int64) (domain.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var event domain.Event

	if event, found := s.Data[id]; found {
		return event, nil
	}

	return event, errors.New("not found in store")
}

func (s *Store) Fetch() ([]domain.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var events []domain.Event

	for _, event := range s.Data {
		events = append(events, event)
	}

	if len(events) < 1 {
		return nil, errors.New("no events in store")
	}

	return events, nil
}

func (s *Store) FetchInDuration(start, end time.Time, userID int64) ([]domain.Event, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	var events []domain.Event

	for _, event := range s.Data {

		if time.Time(event.EventTime) == start || time.Time(event.EventTime) == end || (time.Time(event.EventTime).After(start) && time.Time(event.EventTime).Before(end)) && event.UserID == userID {
			events = append(events, event)
		}
	}

	if len(events) < 1 {
		return nil, errors.New("no events in store")
	}

	return events, nil
}

func (s *Store) Del(id int64) (int64, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.Data, id)
	return id, nil

}
