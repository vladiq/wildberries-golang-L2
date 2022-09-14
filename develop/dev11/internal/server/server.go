package server

import (
	"log"
	"net/http"
	"time"
	cal "wb_l2/develop/dev11/internal/calendar"
)

type calendarService struct {
	address string
	mux     *http.ServeMux
	events  *cal.Calendar
}

func NewCalendarService(address string) *calendarService {
	return &calendarService{address, http.NewServeMux(), cal.NewCalendar()}
}

func logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		log.Printf("%s, %s, %s\n", r.Method, r.URL, time.Since(start))
	}
}

func (s *calendarService) Run() {
	srv := &http.Server{
		Addr:    s.address,
		Handler: s.mux,
	}

	s.mux.HandleFunc("/create_event", logging(s.Create))
	s.mux.HandleFunc("/update_event", logging(s.Update))
	s.mux.HandleFunc("/delete_event", logging(s.Remove))

	s.mux.HandleFunc("/events_for_day", logging(s.EventsDay))
	s.mux.HandleFunc("/events_for_week", logging(s.EventsWeek))
	s.mux.HandleFunc("/events_for_month", logging(s.EventsMonth))

	log.Fatal(srv.ListenAndServe())
}
