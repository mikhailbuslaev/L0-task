package events

import (
	"time"
	"sync"
	"errors"
	"math/rand"
)

var (
	NothingErr = errors.New("nothing to change, empty array")
	WrongIdErr = errors.New("there is not event with that id, try another id")

	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
)

type Event struct {
	Id string `json:"event_id"`
	Title string `json:"title"`
	Time time.Time `json:"time"`
	Description string `json:"description"`
}

type EventPool struct {
	Data map[string][]Event// key of map is user id
	//Eventpool must be alvays sorted by time asc
	sync.RWMutex
}

func NewEventPool() *EventPool {
	evp := EventPool{}
	evp.Data = make(map[string][]Event)
	return &evp
}

func NewEvent(title string, time time.Time, description string) *Event {
	event := Event{Title:title, Time:time, Description:description}
	event.Id = randString()
	return &event
}

func randString() string {
	rand.Seed(time.Now().UnixNano())
    b := make([]byte, 10)
    for i := range b {
        b[i] = letterBytes[rand.Intn(len(letterBytes))]
    }
    return string(b)
}

func NewEvents() []Event {
	events := make([]Event, 0, 10)
	return events
}

func (e *EventPool) EventsDay(actual time.Time, key string) []Event{
	actualYear := actual.Year()
	actualMonth := actual.Month()
	actualDay := actual.Day()
	finded := NewEvents()
	e.RLock()
	defer e.RUnlock()
	for _, v := range e.Data[key] {
		if v.Time.Year() != actualYear {
			return finded
		}
		if v.Time.Month() != actualMonth {
			return finded
		}
		if v.Time.Day() != actualDay {
			return finded
		}
		finded = append(finded, v)
	}
	return finded
}

func (e *EventPool) EventsWeek(actual time.Time, key string) []Event{
	actualYear, actualWeek := actual.ISOWeek()
	finded := NewEvents()
	e.RLock()
	defer e.RUnlock()
	for _, v := range e.Data[key] {
		if v.Time.Year() != actualYear {
			return finded
		}
		if _, week := v.Time.ISOWeek(); week != actualWeek {
			return finded
		}
		finded = append(finded, v)
	}
	return finded
}

func (e *EventPool) EventsMonth(actual time.Time, key string) []Event{
	actualYear := actual.Year()
	actualMonth := actual.Month()
	finded := NewEvents()
	e.RLock()
	defer e.RUnlock()
	for _, v := range e.Data[key] {
		if v.Time.Year() != actualYear {
			return finded
		}
		if v.Time.Month() != actualMonth {
			return finded
		}
		finded = append(finded, v)
	}
	return finded
}

func (evp *EventPool) Add(e *Event, key string) {
	evp.Lock()
	defer evp.Unlock()
	switch length := len(evp.Data[key]); length {
	case 0:
		evp.Data[key] = make([]Event, 0, 10)
		evp.Data[key] = append(evp.Data[key], *e)
	default:
		i := 0
		for i < len(evp.Data[key]){
			if e.Time.Unix() > evp.Data[key][i].Time.Unix() {
				i++
			} else {
				break
			}
		}
		evp.Data[key] = append(evp.Data[key], *e)
		evp.Data[key][i], evp.Data[key][length] = evp.Data[key][length], evp.Data[key][i]
	}
}

func (evp *EventPool) findById(key, id string) int64 {
	for i := range evp.Data[key] {
		if evp.Data[key][i].Id == id {
			return int64(i)
		}
	}
	return -1
}

func (evp *EventPool) Delete(key, id string) error {
	evp.Lock()
	defer evp.Unlock()
	switch length := len(evp.Data[key]); length {
	case 0:
		return NothingErr
	case 1:
		if evp.Data[key][0].Id == id {
			evp.Data[key] = make([]Event, 0, 10)
		} else {
			return WrongIdErr
		}
	default:
		var index int64
		if index = evp.findById(key, id); index == -1 {
			return WrongIdErr
		}
		evp.Data[key] = append(evp.Data[key][:index], evp.Data[key][index+1:]...)
	}
	return nil
}

func (evp *EventPool) Set(key string, e *Event) error {
	evp.Lock()
	defer evp.Unlock()
	switch length := len(evp.Data[key]); length {
	case 0:
		return NothingErr
	default:
		var index int64
		if index = evp.findById(key, e.Id); index == -1 {
			return WrongIdErr
		}
		if e.Title != "" {
			evp.Data[key][index].Title = e.Title
		}
		if e.Time != (time.Time{}) {
			evp.Data[key][index].Time = e.Time
		}
		if e.Description != "" {
			evp.Data[key][index].Description = e.Description
		}
	}
	return nil
}


