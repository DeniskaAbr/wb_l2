package domain

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type Event struct {
	ID          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	EventTime   JSONTime  `json:"event_time"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	UserID      int64     `json:"user_id"`
}

type DeleteEvent struct {
	ID int64 `json:"id"`
}

type JSONTime time.Time

func (j *JSONTime) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	fmt.Println(s)
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}
	*j = JSONTime(t)
	return nil
}

func (j JSONTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(j))
}

func (j JSONTime) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
