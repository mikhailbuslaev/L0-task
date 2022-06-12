package subscriber

import (
	"testing"
	"nats-subscriber/cache"
)

func Test_CheckDB(t *testing.T) {
	s := [3]*Subscriber{}
	got := [3]bool{}
	for i := range s {
		s[i] = New()
	}
	//1 case is normal
	s[1].DbConfig.User = "fdsgsdg"//2 case, wrong user
	s[2].DbConfig.Host = "asdsd8881"//3 case, wrong host

	for i := range s {
		got[i] = s[i].checkDB()
	}
    if got[0] != true {
        t.Errorf("Got error, dont want it")
    }
	if got[1] == true {
        t.Errorf("Want error, got nothing")
    }
	if got[2] == true {
        t.Errorf("Want error, got nothing")
    }
}

func Test_connectToDB(t *testing.T) {
	s := [3]*Subscriber{}
	got := [3]error{}
	for i := range s {
		s[i] = New()
	}
	//1 case is normal
	s[1].DbConfig.User = "fdsgsdg"//2 case, wrong user
	s[2].DbConfig.Host = "asdsd8881"//3 case, wrong host

	for i := range s {
		_, got[i] = s[i].connectToDB()
	}
    if got[0] != nil {
        t.Errorf("Got error, dont want it")
    }
	if got[1] == nil {
        t.Errorf("Want error, got nothing")
    }
	if got[2] == nil {
        t.Errorf("Want error, got nothing")
    }
}

func Test_restoreCache(t *testing.T) {
	s := [3]*Subscriber{}
	got := [3]error{}
	for i := range s {
		s[i] = New()
	}
	//1 case is normal
	s[1].DbConfig.User = "fdsgsdg"//2 case, wrong user
	s[2].DbConfig.Host = "asdsd8881"//3 case, wrong host

	for i := range s {
		db, _ := s[i].connectToDB()
//		defer db.Close()
		got[i] = s[i].restoreCache(db)
	}
    if got[0] != nil {
        t.Errorf("Got error, dont want it")
    }
	if got[1] == nil {
        t.Errorf("Want error, got nothing")
    }
	if got[2] == nil {
        t.Errorf("Want error, got nothing")
    }
}

func Test_pushToDB(t *testing.T) {
	s := New()
	msg := [3]cache.Order{}
	got := [3]error{}
	msg[1].Id = "AFASFA"
	msg[1].Data = `{"aadad":`

	msg[2].Id = "aaqqqqq"
	msg[2].Data = `{"aadad":"dsfsdfsdf"}`
	//1 case is empty message, we want error
	//2 case have invalid json, we want error
	//3 case correct

	for i := range msg {
		db, _ := s.connectToDB()
		defer db.Close()
		got[i] = s.pushToDB(msg[i], db)
	}
    if got[0] == nil {
        t.Errorf("Want error, got nothing")
    }
	if got[1] == nil {
        t.Errorf("Want error, got nothing")
    }
	if got[2] != nil {
        t.Errorf("Got error, dont want it")
    }
}