package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	go runProcess("Process 1", 20)
	go runProcess("Process 2", 20)

	var s string
	fmt.Scanln(&s)

}

func runProcess(name string, total int) {
	for i := 0; i < total; i++ {
		t := time.Duration(rand.Intn(255))
		time.Sleep(time.Millisecond * t)
		fmt.Println(name, i)
	}
}
