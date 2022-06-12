package cache

import (
	"testing"
)

func Test_Add(t *testing.T) {
	want := map[string]Order{"test_id":Order{Id:"test_id"}} 
	c := New()
	c.Add(&Order{Id:"test_id"})
	if c.Data["test_id"] != want["test_id"] {
		t.Errorf("Want not that cache")
	}
}

func Test_Get(t *testing.T) {
	//1 case, we have needed record
	//2 case, we dont have needed record
	want := [2]Order{{Id:"test_id"},{}}
	c := New()
	c.Add(&Order{Id:"test_id"})
	got1, _ := c.Get("test_id")
	got2, _ := c.Get("???")
	if got1 != want[0] {
		t.Errorf("Dont want it")
	}
	if got2 != want[1] {
		t.Errorf("Dont want it")
	}
}