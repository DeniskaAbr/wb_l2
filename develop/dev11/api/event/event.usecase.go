package event

import (
	"dev11/api/domain"
	"log"
	"time"
)

type eventUsecase struct {
	eventRepo domain.IEventRepository
	logger    *log.Logger
}

func NewEventUsecase(ev domain.IEventRepository, logger *log.Logger) *eventUsecase {
	var uc = eventUsecase{eventRepo: ev, logger: logger}

	return &uc
}

func (e *eventUsecase) Fetch() (*[]domain.Event, error) {
	events, err := e.eventRepo.RepoFetch()
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (e *eventUsecase) FetchForDuration(interval domain.TimeInterval, t time.Time, userID int64) (*[]domain.Event, error) {
	events, err := e.eventRepo.RepoFetchForDuration(interval, t, userID)
	if err != nil {
		return nil, err
	}

	return events, nil
}

func (e *eventUsecase) GetByID(id int64) (*domain.Event, error) {
	event, err := e.eventRepo.RepoGetByID(id)
	if err != nil {
		return &domain.Event{}, err
	}

	return event, nil
}

func (e *eventUsecase) Create(event *domain.Event) (int64, error) {

	return e.eventRepo.RepoCreate(event)
}

func (e *eventUsecase) Update(event *domain.Event) (int64, error) {
	updatedID, err := e.eventRepo.RepoUpdate(event)
	if err != nil {
		return 0, err
	}
	return updatedID, nil
}

func (e *eventUsecase) Delete(id int64) (int64, error) {
	existed, _ := e.eventRepo.RepoGetByID(id)
	if *existed == (domain.Event{}) {
		return 0, domain.ErrNotFound
	}

	_, err := e.eventRepo.RepoDelete(id)

	return existed.ID, err

}
