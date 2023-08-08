package belajar_goroutines

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

//membuat channel

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

// channel sebagai parameter

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)

	channel <- "cino"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
	close(channel)
}

// channel in and out
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)

	channel <- "cino"
}

func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

// membuat buffered channel
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 1)

	channel <- "cino"

	defer close(channel)

}

// range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	}()
	for data := range channel {
		fmt.Println("menerima data", data)
	}
	fmt.Println("done")
}

//select channel

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter == 2 {
			break
		}
	}
}

//select channel

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Data dari channel 1 ", data)
			counter++
		case data := <-channel2:
			fmt.Println("Data dari channel 2 ", data)
			counter++

		}
		if counter == 2 {
			break
		}
	}
}
