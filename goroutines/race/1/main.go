package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var result int
var waitGroup sync.WaitGroup
var m sync.Mutex

func main() {

	waitGroup.Add(2)

	go runProcess("Process 1", 20)
	go runProcess("Process 2", 20)

	waitGroup.Wait()
	fmt.Println("Result:", result)
}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		m.Lock()
		result++
		fmt.Println(name, "->", i, "total", result)
		m.Unlock()
	}
	waitGroup.Done()
}
