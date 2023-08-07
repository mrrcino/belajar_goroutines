package belajar_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "cino"
		fmt.Println("Selesai mengambil data")
	}()

	data := <-channel
	fmt.Println(data)
	close(channel)
}
