package main

// We cant use map with concurrent writing (https://go.dev/doc/faq#atomic_maps)
// but we can make read/write queue with mutex
import (
	"fmt"
	"os"
	"sync"
)

type Map struct {
	sync.RWMutex
	Data map[string]string
}

func read(m *Map, key string) {
	m.RLock()
	defer m.RUnlock()
	fmt.Fprintf(os.Stdout, "Read value: %s from map[%s]...\n", m.Data[key], key)
}

func write(m *Map, key, value string) {
	m.Lock()
	defer m.Unlock()
	fmt.Fprintf(os.Stdout, "Write to map[%s] value %s...\n", key, value)
	m.Data[key] = value
}

func main() {
	m := &Map{}
	m.Data = make(map[string]string)
	// Concurent writing
	wg := sync.WaitGroup{}
	wg.Add(3)
	go func() {
		write(m, "key1", "value1")
		wg.Done()
	}()

	go func() {
		write(m, "key2", "value2")
		wg.Done()
	}()

	go func() {
		write(m, "key3", "value3")
		wg.Done()
	}()
	wg.Wait()
	//Concurent reading
	wg.Add(3)
	go func() {
		read(m, "key1")
		wg.Done()
	}()
	go func() {
		read(m, "key2")
		wg.Done()
	}()
	go func() {
		read(m, "key3")
		wg.Done()
	}()
	wg.Wait()
}
