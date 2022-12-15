package domain

import "time"

const (
	Day TimeInterval = iota
	Week
	Month
	Year
)

type TimeInterval int

type IEventUsecase interface {
	Fetch() (events *[]Event, err error)
	FetchForDuration(interval TimeInterval, t time.Time, userID int64) (events *[]Event, err error)
	GetByID(id int64) (event *Event, err error)

	Create(event *Event) (eventID int64, err error)
	Update(event *Event) (updatedID int64, err error)
	Delete(id int64) (int64, error)
}

type IEventRepository interface {
	RepoFetch() (events *[]Event, err error)
	RepoFetchForDuration(interval TimeInterval, t time.Time, userID int64) (events *[]Event, err error)
	RepoGetByID(id int64) (event *Event, err error)

	RepoCreate(event *Event) (eventID int64, err error)
	RepoUpdate(event *Event) (updatedID int64, err error)
	RepoDelete(id int64) (int64, error)
}

//https://stackoverflow.com/questions/36830212/get-the-first-and-last-day-of-current-month-in-go-golang
