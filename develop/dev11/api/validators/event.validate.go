package validators

import (
	"dev11/api/domain"
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"time"
)

func EventValidate(r *http.Request) (*domain.Event, error) {
	var event domain.Event
	err := json.NewDecoder(r.Body).Decode(&event)
	return &event, err
}

func EventsValidate(r *http.Request) (*[]domain.Event, error) {
	var events []domain.Event
	err := json.NewDecoder(r.Body).Decode(&events)
	return &events, err
}

func DeleteEventValidate(r *http.Request) (*domain.DeleteEvent, error) {
	var delEvent domain.DeleteEvent
	err := json.NewDecoder(r.Body).Decode(&delEvent)
	return &delEvent, err
}

func IntervalEventHandlerValidate(r *http.Request) (int64, time.Time, error) {

	r.ParseForm()

	user_id := r.URL.Query().Get("user_id")
	date := r.URL.Query().Get("date")

	// fmt.Println("request.Form::")
	// for key, value := range r.Form {
	// 	fmt.Printf("Key:%s, Value:%s\n", key, value)
	// }
	// fmt.Println("\nrequest.PostForm::")
	// for key, value := range r.PostForm {
	// 	fmt.Printf("Key:%s, Value:%s\n", key, value)
	// }

	if user_id != "" && date != "" {
		userID, err := strconv.Atoi(user_id)

		if err != nil {
			return 0, time.Now(), err
		}

		t, err := time.Parse("2006-01-02", date)
		if err != nil {
			return 0, time.Now(), err
		}

		return int64(userID), t, nil
	}

	return 0, time.Now(), errors.New("not descripted validate error")
}
