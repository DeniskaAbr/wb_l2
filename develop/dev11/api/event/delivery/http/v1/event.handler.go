package v1

import (
	"dev11/api/domain"
	"dev11/api/validators"
	"encoding/json"
	"errors"
	"net/http"
)

func (eh *EventHandler) CreateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseWithError(w, domain.ErrMethodNotAllowed, http.StatusMethodNotAllowed, "")
		return
	}

	// body, error := ioutil.ReadAll(r.Body)
	// if error != nil {
	// 	fmt.Println(error)
	// }

	// fmt.Println(string(body))

	event, err := validators.EventValidate(r)

	if err != nil {
		ResponseWithError(w, domain.ErrBadParamInput, http.StatusBadRequest, err.Error())
		return
	}
	event.ID = -1

	id, err := eh.uc.Create(event)

	if err != nil {
		ResponseWithError(w, domain.ErrInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}
	event.ID = id

	ResponseWithResult(w, []domain.Event{*event})
}

func (eh *EventHandler) UpdateEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseWithError(w, domain.ErrMethodNotAllowed, http.StatusMethodNotAllowed, "")
		return
	}

	event, err := validators.EventValidate(r)

	if err != nil {
		ResponseWithError(w, domain.ErrBadParamInput, http.StatusBadRequest, err.Error())
		return
	}

	_, err = eh.uc.Update(event)

	if err != nil {
		ResponseWithError(w, domain.ErrInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithResult(w, []domain.Event{*event})
}

func (eh *EventHandler) DeleteEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		ResponseWithError(w, domain.ErrMethodNotAllowed, http.StatusMethodNotAllowed, "")
		return
	}

	delEvent, err := validators.DeleteEventValidate(r)

	if err != nil {
		ResponseWithError(w, domain.ErrBadParamInput, http.StatusBadRequest, err.Error())
		return
	}

	id, err := eh.uc.Delete(delEvent.ID)

	if err != nil {
		ResponseWithError(w, domain.ErrInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	DeleteResponseWithResult(w, id)
}

func (eh *EventHandler) ForDayEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseWithError(w, domain.ErrMethodNotAllowed, http.StatusMethodNotAllowed, "")
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		ResponseWithError(w, errors.New(""), http.StatusUnsupportedMediaType, "")
		return
	}

	userID, t, err := validators.IntervalEventHandlerValidate(r)

	if err != nil {
		ResponseWithError(w, domain.ErrBadParamInput, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eh.uc.FetchForDuration(domain.Day, t, userID)

	if err != nil {
		ResponseWithError(w, domain.ErrInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithResult(w, *events)
}

func (eh *EventHandler) ForWeekEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseWithError(w, domain.ErrMethodNotAllowed, http.StatusMethodNotAllowed, "")
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		ResponseWithError(w, errors.New(""), http.StatusUnsupportedMediaType, "")
		return
	}

	userID, t, err := validators.IntervalEventHandlerValidate(r)

	if err != nil {
		ResponseWithError(w, domain.ErrBadParamInput, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eh.uc.FetchForDuration(domain.Week, t, userID)

	if err != nil {
		ResponseWithError(w, domain.ErrInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithResult(w, *events)
}

func (eh *EventHandler) ForMonthEventHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ResponseWithError(w, domain.ErrMethodNotAllowed, http.StatusMethodNotAllowed, "")
		return
	}

	headerContentTtype := r.Header.Get("Content-Type")
	if headerContentTtype != "application/x-www-form-urlencoded" {
		ResponseWithError(w, errors.New(""), http.StatusUnsupportedMediaType, "")
		return
	}

	userID, t, err := validators.IntervalEventHandlerValidate(r)

	if err != nil {
		ResponseWithError(w, domain.ErrBadParamInput, http.StatusBadRequest, err.Error())
		return
	}

	events, err := eh.uc.FetchForDuration(domain.Month, t, userID)

	if err != nil {
		ResponseWithError(w, domain.ErrInternalServerError, http.StatusInternalServerError, err.Error())
		return
	}

	ResponseWithResult(w, *events)
}

type ResultResponse struct {
	Events []domain.Event `json:"result"`
}

func ResponseWithResult(w http.ResponseWriter, events []domain.Event) error {
	resp, err := json.Marshal(ResultResponse{events})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}

type DeleteResponse struct {
	ID int64 `json:"id"`
}

func DeleteResponseWithResult(w http.ResponseWriter, id int64) error {
	resp, err := json.Marshal(DeleteResponse{ID: id})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return err
	}

	_, err = w.Write(resp)
	if err != nil {
		return err
	}
	return nil
}

type ErrorResponse struct {
	Err    string `json:"error,omitempty"`
	Detail string `json:"detail,omitempty"`
}

func ResponseWithError(w http.ResponseWriter, err error, code int, detail string) error {
	resp, marshalError := json.Marshal(ErrorResponse{err.Error(), detail})
	if marshalError != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return marshalError
	}

	http.Error(w, string(resp), code)
	return nil
}
