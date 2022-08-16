package main

import (
	"fmt"
	"math/rand"
)

func producer(itemCh chan<- int) {
	rand.Seed(26)
	for i := 0; i < 10; i++ {
		n := rand.Intn(10)
		itemCh <- n
	}
	close(itemCh)
}

func consume(itemCh <-chan int, join chan<- int) {
	for {
		n, ok := <-itemCh
		if !ok {
			join <- 1
			break
		}
		if n%2 == 0 {
			fmt.Println("O número ", n, " é par")
		} else {
			fmt.Println("O número ", n, " é impar")
		}
	}
}

func main() {
	itemCh := make(chan int)
	join := make(chan int)

	go producer(itemCh)
	go consume(itemCh, join)

	<-join
}
