package main

import (
	"fmt"
	"sync"
)

func main() {
	karta := sync.Map{}
	wg := sync.WaitGroup{}
	// Concurrent write to map
	wg.Add(3)
	go func() {
		karta.Store("key1", "value1")
		wg.Done()
	}()

	go func() {
		karta.Store("key2", "value2")
		wg.Done()
	}()

	go func() {
		karta.Store("key3", "value3")
		wg.Done()
	}()
	wg.Wait()
	// Concurrent read map
	wg.Add(3)
	go func() {
		value, _ := karta.Load("key1")
		fmt.Println(value)
		wg.Done()
	}()

	go func() {
		value, _ := karta.Load("key2")
		fmt.Println(value)
		wg.Done()
	}()

	go func() {
		value, _ := karta.Load("key3")
		fmt.Println(value)
		wg.Done()
	}()
	wg.Wait()
}
