package server

import (
	"encoding/json"
	"github.com/google/uuid"
	"net/http"
	"time"
)

func (s *calendarService) Create(w http.ResponseWriter, r *http.Request) {
	if validateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		} else if !ev.validate() {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, "validation failed")
		} else {
			ev.ID = uuid.New().String()
			s.events.AddEventToCalendar(ev.modelToEvent())
			sendJsonResponse(false, w, http.StatusOK, "event created successfully")
		}
	}
}

func (s *calendarService) Update(w http.ResponseWriter, r *http.Request) {
	if validateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		} else if err := s.events.UpdateEvent(ev.modelToEvent()); err != nil {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		} else {
			sendJsonResponse(false, w, http.StatusOK, "event updated successfully")
		}
	}
}

func (s *calendarService) Remove(w http.ResponseWriter, r *http.Request) {
	if validateQuery(w, r, http.MethodPost) {
		ev := EventModel{}
		if err := json.NewDecoder(r.Body).Decode(&ev); err != nil || ev.ID == "" {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, "decoding error")
			return
		}
		if err := s.events.DeleteEvent(ev.ID); !err {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, "cannot find event by ID")
			return
		}
		sendJsonResponse(false, w, http.StatusOK, "event deleted successfully")
	}
}

func (s *calendarService) EventsDay(w http.ResponseWriter, r *http.Request) {
	s.eventsInterval(w, r, 0, 0, 1)
}

func (s *calendarService) EventsWeek(w http.ResponseWriter, r *http.Request) {
	s.eventsInterval(w, r, 0, 0, 7)
}

func (s *calendarService) EventsMonth(w http.ResponseWriter, r *http.Request) {
	s.eventsInterval(w, r, 0, 1, 0)
}

func (s *calendarService) eventsInterval(w http.ResponseWriter, r *http.Request, years, months, days int) {
	if validateQuery(w, r, http.MethodGet, "user_id", "date") {
		if date, err := time.Parse("2006-01-02", r.URL.Query().Get("date")); err != nil {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, err.Error())
		} else {
			userID := r.URL.Query().Get("user_id")
			evs := s.events.GetEventsInTimeRange(userID, date, date.AddDate(years, months, days))
			sendJsonResponse(false, w, http.StatusOK, evs)
		}
	}
}

type resultMsg struct {
	Msg interface{} `json:"result"`
}

type errorMsg struct {
	Msg string `json:"error"`
}

func sendJsonResponse(err bool, w http.ResponseWriter, code int, msg interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(code)
	if err {
		errResp := errorMsg{msg.(string)}
		if err := json.NewEncoder(w).Encode(errResp); err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
		}
	} else {
		resResp := resultMsg{msg}
		if err := json.NewEncoder(w).Encode(resResp); err != nil {
			http.Error(w, "error", http.StatusInternalServerError)
		}
	}
}

func validateQuery(w http.ResponseWriter, r *http.Request, query ...string) bool {
	if r.Method != query[0] {
		sendJsonResponse(true, w, http.StatusMethodNotAllowed, "bad request method")
		return false
	}
	for _, v := range query[1:] {
		if !r.URL.Query().Has(v) {
			sendJsonResponse(true, w, http.StatusServiceUnavailable, "bad parameters")
			return false
		}
	}
	return true
}
