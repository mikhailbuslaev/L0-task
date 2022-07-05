package server

import (
	"net/http"
	"github.com/mikhailbuslaev/wb-tasks/l2/dev11/events"
	"time"
	"os"
	"encoding/json"
	"strconv"
)

type Response struct {
	Result interface{} `json:"result"`
}

type Server struct {
	Port string
	LogFileName string
	LogFile *os.File
}

func (s *Server) Logger(b []byte) error{
	currentTime := time.Now()
	byteTime, err := currentTime.MarshalText()
	if err != nil {
		return err
	}
	if _, err := s.LogFile.Write(byteTime); err != nil {
		return err
	}
	if _, err := s.LogFile.Write([]byte("\n")); err != nil {
		return err
	}
	if _, err := s.LogFile.Write(b); err != nil {
		return err
	}
	if _, err := s.LogFile.Write([]byte("\n\n")); err != nil {
		return err
	}
	return nil
}

func (s *Server) Response(w http.ResponseWriter, data []byte, status int) {
	w.Write(data)
	w.WriteHeader(status)
	if err := s.Logger(data); err != nil {
		err.Error()
	}
}

func (s *Server) ListenAndServe() {
	evp := events.NewEventPool()
	var err error
	s.LogFile, err = os.OpenFile(s.LogFileName, os.O_APPEND, 0755)
	if err != nil {
		err.Error()
		os.Exit(1)
	}
	defer s.LogFile.Close()

	getEvent := func(w http.ResponseWriter, req *http.Request, mode string) {
		if req.Method != "GET" {
			s.Response(w, []byte("error: server accept only GET method at this address"), http.StatusMethodNotAllowed)
			return
		}
		userId := req.FormValue("user_id")
		if userId == "" {
			s.Response(w, []byte("error: user_id param missing"), http.StatusBadRequest)
			return
		}
		currentTime := time.Now()
		events := events.NewEvents()
		switch mode {
		case "for_day":
			events = evp.EventsDay(currentTime, userId)
		case "for_week":
			events = evp.EventsWeek(currentTime, userId)
		case "for_month":
			events = evp.EventsMonth(currentTime, userId)
		}
		
		resp := Response{Result:events}

		b, err := json.Marshal(resp)
		if err != nil {
			s.Response(w, []byte("error: server cannot make response for you"), http.StatusInternalServerError)
			return
		}
		s.Response(w, b, http.StatusOK)
	}

	eventsForDay := func(w http.ResponseWriter, req *http.Request) {getEvent(w, req, "for_day")}
	eventsForWeek := func(w http.ResponseWriter, req *http.Request) {getEvent(w, req, "for_week")}
	eventsForMonth := func(w http.ResponseWriter, req *http.Request) {getEvent(w, req, "for_month")}

	eventCreate := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			s.Response(w, []byte("error: server accept only POST method at this address"), http.StatusMethodNotAllowed)
			return
		}
		userId := req.FormValue("user_id")
		if userId == "" {
			s.Response(w, []byte("error: user_id param missing"), http.StatusBadRequest)
			return
		}
		timeString := req.FormValue("time")
		if timeString == "" {
			s.Response(w, []byte("error: time param missing"), http.StatusBadRequest)
			return
		}
		title := req.FormValue("title")
		if title == "" {
			s.Response(w, []byte("error: title param missing"), http.StatusBadRequest)
			return
		}
		description := req.FormValue("description")
		intTime, err := strconv.Atoi(timeString)
		if err != nil {
			s.Response(w, []byte("error: server cannot parse your date"), http.StatusBadRequest)
			return
		}
		time := time.Unix(int64(intTime), 0).UTC()
		event := events.NewEvent(title, time, description)
		evp.Add(event, userId)

		resp := Response{Result:"event created succesfully"}

		b, err := json.Marshal(resp)
		if err != nil {
			s.Response(w, []byte("error: server cannot make response for you"), http.StatusInternalServerError)
			return
		}
		s.Response(w, b, http.StatusOK)
	}

	eventUpdate := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			s.Response(w, []byte("error: server accept only POST method at this address"), http.StatusMethodNotAllowed)
			return
		}
		event := &events.Event{}
		userId := req.FormValue("user_id")
		if userId == "" {
			s.Response(w, []byte("error: user_id param missing"), http.StatusBadRequest)
			return
		}
		event.Id = req.FormValue("event_id")
		if event.Id == "" {
			s.Response(w, []byte("error: event_id param missing"), http.StatusBadRequest)
			return
		}
		timeString := req.FormValue("time")
		var err error
		intTime, err := strconv.Atoi(timeString)
		if err != nil {
			s.Response(w, []byte("error: server cannot parse your date"), http.StatusBadRequest)
			return
		}

		event.Time = time.Unix(int64(intTime), 0).UTC()
		event.Title = req.FormValue("title")
		event.Description = req.FormValue("description")

		if err = evp.Set(userId, event); err != nil {
			s.Response(w, []byte("error: server cannot update event"), http.StatusInternalServerError)
			return
		}

		resp := Response{Result:"event update succesfully"}

		b, err := json.Marshal(resp)
		if err != nil {
			s.Response(w, []byte("error: server cannot make response for you"), http.StatusInternalServerError)
			return
		}
		s.Response(w, b, http.StatusOK)
	}

	eventDelete := func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			s.Response(w, []byte("error: server accept only POST method at this address"), http.StatusMethodNotAllowed)
			return
		}
		userId := req.FormValue("user_id")
		if userId == "" {
			s.Response(w, []byte("error: user_id param missing"), http.StatusBadRequest)
			return
		}
		eventId := req.FormValue("event_id")
		if eventId == "" {
			s.Response(w, []byte("error: event_id param missing"), http.StatusBadRequest)
			return
		}

		if err := evp.Delete(userId, eventId); err != nil {
			s.Response(w, []byte("error: server cannot delete event"), http.StatusInternalServerError)
			return
		}

		resp := Response{Result:"event delete succesfully"}

		b, err := json.Marshal(resp)
		if err != nil {
			s.Response(w, []byte("error: server cannot make response for you"), http.StatusInternalServerError)
			return
		}
		s.Response(w, b, http.StatusOK)
	}

	http.HandleFunc("/events_for_day", eventsForDay)
	http.HandleFunc("/events_for_week", eventsForWeek)
	http.HandleFunc("/events_for_month", eventsForMonth)

	http.HandleFunc("/create_event", eventCreate)
	http.HandleFunc("/update_event", eventUpdate)
	http.HandleFunc("/delete_event", eventDelete)
		
	http.ListenAndServe(s.Port, nil)
}