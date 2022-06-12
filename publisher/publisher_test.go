package publisher

import (
	"testing"
	stan "github.com/nats-io/stan.go"
)

func Test_Publish(t *testing.T) {
	//1 case, correct publisher
	//2 case, wrong cluster
	//3 case, wrong channel
	want := [3]string{"test_message","",""}
	var got [3]string
	pub := [3]*Publisher{New(), New(), New()}
	pub[1].Cluster = "wrong-cluster"
	pub[2].Channel = "wrong-channel"
	
	for i:= range want {
		sc, _ := stan.Connect("test-cluster", "sub")
		sub, _ := sc.Subscribe("foo", func(m *stan.Msg) {
			got[i] = string(m.Data)
		})
		pub[i].Publish([]byte("test_message"))//1 case, correct
	
		sub.Unsubscribe()
		sc.Close()
	
		if got[i] != want[i] {
			t.Errorf("Dont want it")
		}
	}
}