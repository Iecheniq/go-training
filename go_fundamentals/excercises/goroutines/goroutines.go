package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string, 1)
	end := make(chan bool, 1)
	closeChannel := false

	go func() {
		for {
			func() {
				time.Sleep(time.Second)
				message <- "Reading"
			}()

			fmt.Println(<-message)

			func(bool) {
				if closeChannel {
					fmt.Printf("Channel is closed")
					close(message)
					close(end)
					return
				}
			}(closeChannel)
		}
	}()

	time.Sleep(5 * time.Second)
	closeChannel = true
	<-end

}
