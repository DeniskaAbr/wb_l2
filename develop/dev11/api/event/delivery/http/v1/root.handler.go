package v1

import (
	"dev11/api/domain"
	"dev11/api/middleware"
	"fmt"
	"log"
	"net/http"
)

const (
	API_PATH = "/api/v1"
)

type EventHandler struct {
	Mux   *http.ServeMux
	loger *log.Logger
	uc    domain.IEventUsecase
}

func NewEventHandler(uc domain.IEventUsecase, logger *log.Logger) *EventHandler {
	mux := http.NewServeMux()

	eh := EventHandler{
		Mux:   mux,
		loger: logger,
		uc:    uc,
	}

	eh.AddHandlers()

	return &eh
}

func (eh *EventHandler) AddHandlers() {
	eh.Mux.HandleFunc("/", eh.handler)

	eh.Mux.HandleFunc(API_PATH+"/create_event", middleware.Logging(eh.CreateEventHandler, eh.loger))
	eh.Mux.HandleFunc(API_PATH+"/update_event", middleware.Logging(eh.UpdateEventHandler, eh.loger))
	eh.Mux.HandleFunc(API_PATH+"/delete_event", middleware.Logging(eh.DeleteEventHandler, eh.loger))
	eh.Mux.HandleFunc(API_PATH+"/events_for_day", middleware.Logging(eh.ForDayEventHandler, eh.loger))
	eh.Mux.HandleFunc(API_PATH+"/events_for_week", middleware.Logging(eh.ForWeekEventHandler, eh.loger))
	eh.Mux.HandleFunc(API_PATH+"/events_for_month", middleware.Logging(eh.ForMonthEventHandler, eh.loger))
}

// POST "/create_event"
// POST "/update_event"
// POST "/delete_event"
// GET "/events_for_day"
// GET "/events_for_week"
// GET "/events_for_month"

func (eh *EventHandler) handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		// http.NotFound(w, r)
		return
	}

	fmt.Fprintf(w, "Hello World")
}
