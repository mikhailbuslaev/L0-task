package subscriber

import (
	"nats-subscriber/cache"
	"testing"
)

func Test_CheckDB(t *testing.T) {
	s := [3]*Subscriber{}
	got := [3]bool{}
	for i := range s {
		s[i] = New()
	}
	//1 case is normal
	s[1].DbConfig.User = "fdsgsdg"   //2 case, wrong user
	s[2].DbConfig.Host = "asdsd8881" //3 case, wrong host

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
	s[1].DbConfig.User = "fdsgsdg"   //2 case, wrong user
	s[2].DbConfig.Host = "asdsd8881" //3 case, wrong host

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
	s[1].DbConfig.User = "fdsgsdg"   //2 case, wrong user
	s[2].DbConfig.Host = "asdsd8881" //3 case, wrong host

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

	msg[2].Id = "aaqa=d9pp"
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

func Test_pushToFile(t *testing.T) {
	//1 case, pushing correct order, waitnig no errors
	//2 case, pushing empty order, waiting no errors
	sub := New()
	sub.RestoreFile = "subscriber_test_ok.csv"
	got := [2]error{}
	orders := [2]*cache.Order{{Id: "", Data: ""}, {Id: "some_id_1", Data: `{"data":"some_data"}`}}
	for i := range got {
		got[i] = sub.pushToFile(orders[i])
		if got[i] != nil {
			t.Errorf("Got error, dont want it")
		}
	}
}

func Test_restoreFromFile(t *testing.T) {
	//1 case, restoring from correct file with ok records, waiting no errors
	//2 case, restoring empty file, waiting no errors
	//3 case, restoring fron non-existing file, waiting error
	got := [3]error{}
	subs := [3]*Subscriber{New(), New(), New()}
	for i := range subs {
		subs[i].Cache = cache.New()
	}
	subs[0].RestoreFile = "subscriber_test_ok.csv"
	subs[1].RestoreFile = "subscriber_test_empty.csv"
	subs[2].RestoreFile = "subscriber_test_non_existing_file.csv"
	for i := range subs {
		got[i] = subs[i].restoreFromFile()
	}
	if got[0] != nil {
		t.Errorf("Got error, dont want it")
	}
	if got[1] != nil {
		t.Errorf("Got error, dont want it")
	}
	if got[2] == nil {
		t.Errorf("Want error, got nothing")
	}
}
