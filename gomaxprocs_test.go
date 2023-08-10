package belajar_goroutines

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestGetGoMaxProcs(t *testing.T) {
	group := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			time.Sleep(3 * time.Second)
			group.Done()
		}()
	}
	totalcpu := runtime.NumCPU()
	fmt.Println("total CPU", totalcpu)

	totalThread := runtime.GOMAXPROCS(-1)
	fmt.Println("total thread", totalThread)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("total Go ROutine", totalGoRoutine)

}
