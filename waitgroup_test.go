package belajar_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

func RunAsync(group *sync.WaitGroup) {
	defer group.Done()

	group.Add(1)
	fmt.Println("hello")

}
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group)
	}

	group.Wait()
	fmt.Println("selesai")
}
