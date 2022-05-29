package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var waitGroup sync.WaitGroup

func main() {
	waitGroup.Add(2)

	go runProcess("Process 1", 20)
	go runProcess("Process 2", 20)

	waitGroup.Wait()

}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		fmt.Println(name, i)
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
	}
	waitGroup.Done()
}
