package event

import (
	"dev11/api/domain"
	"dev11/internal/utils"
	"log"
	"time"
)

// inmemEventRepository
type inmemEventRepository struct {
	store  *Store
	logger *log.Logger
}

// NewMemoryEventRepository
func NewMemoryEventRepository(store *Store, logger *log.Logger) *inmemEventRepository {
	var imr = inmemEventRepository{store: store, logger: logger}
	return &imr
}

// RepoFetch
func (m *inmemEventRepository) RepoFetch() (*[]domain.Event, error) {

	events, err := m.store.Fetch()
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// RepoFetchForDuration
func (m *inmemEventRepository) RepoFetchForDuration(interval domain.TimeInterval, t time.Time, userID int64) (*[]domain.Event, error) {
	var err error
	var events = make([]domain.Event, 0)

	start, end, err := utils.StartEndFromTimeInterval(interval, t)

	if err != nil {
		return nil, err
	}

	events, err = m.store.FetchInDuration(start, end, userID)
	if err != nil {
		return nil, err
	}

	return &events, nil
}

// RepoGetByID
func (m *inmemEventRepository) RepoGetByID(id int64) (*domain.Event, error) {
	var event domain.Event

	event, err := m.store.GetById(id)
	if err != nil {
		return &domain.Event{}, err
	}

	return &event, nil
}

// RepoCreate
func (m *inmemEventRepository) RepoCreate(event *domain.Event) (int64, error) {
	id := m.store.Put(event)

	return id, nil
}

// RepoUpdate
func (m *inmemEventRepository) RepoUpdate(event *domain.Event) (int64, error) {
	_, err := m.RepoGetByID(event.ID)

	if err != nil {
		return 0, err
	}

	id := m.store.Put(event)

	return id, nil
}

// RepoDelete
func (m *inmemEventRepository) RepoDelete(id int64) (int64, error) {
	deletedID, err := m.store.Del(id)

	return deletedID, err

}
