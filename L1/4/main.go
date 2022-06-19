package main
import (
	"fmt"
	"os"
	"os/signal"
	"math/rand"
	"time"
)
// Worker function
func listen(dataChannel chan int, id int, quitChannel chan bool) {
	for {
		// Worker trying to catch quit signal or data message
		select {
		case <- quitChannel:
			fmt.Fprintf(os.Stdout, "Quit worker '%d'.\n", id)
			return
		case <- dataChannel :
			fmt.Fprintf(os.Stdout, "Worker '%d' receive message '%d'.\n", id, <-dataChannel)	
		default:
		}
	}
}
// Poster function
func post(dataChannel chan int, quitChannel chan bool) {
	for {
		// Poster trying to catch quit signal, by default posts messages to channel
		select {
		case <- quitChannel:
			fmt.Println("Quit publisher")
			return
		default:
			time.Sleep(400*time.Millisecond)
			dataChannel <- rand.Int()
		}
	}
}

func main() {
	// This channel for data exchange
	dataChannel := make(chan int)
	// This channel for handling safety quit
	quitChannel := make(chan bool)
	// This channel catches CtrC interrupt
	signalChannel := make(chan os.Signal, 1)
	var n int
	fmt.Print("Enter num of workers: ")
	fmt.Scan(&n)
	for i:=0; i<n;i++ {
		// Running workers
		go listen(dataChannel, i, quitChannel)
	}
	// Running posting
	go post(dataChannel, quitChannel)

	signal.Notify(signalChannel, os.Interrupt)
	// Block main routine until CtrC signal
	<-signalChannel
	// Closing starts here
	fmt.Println("Received an interrupt, closing program...")
	// 6 Closing signals, 5 for workers, 1 for poster
	for i:=0; i<=n;i++ {
		quitChannel <- true
	}
	time.Sleep(3*time.Second)
}