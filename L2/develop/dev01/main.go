package main

import (
	"github.com/beevik/ntp"
	"os"
	"time"
	"fmt"
)

func main() {
	response, err := ntp.Query("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println("Current time according to server is: ", response.Time)
	fmt.Println("Guessed 'actual' time is: ", time.Now().Add(response.ClockOffset))
}