package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	join := make(chan int)
	n := 5

	for i := 0; i < n; i++ {
		rand.Seed(12)
		go func(i int) {
			sleepTime := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(sleepTime)
			fmt.Println("hi im:  ", i)
			join <- 1
		}(i)
	}

	for i := 0; i < n; i++ {
		<-join
	}

	fmt.Println("done")
}
