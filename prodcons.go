package main

import (
	"fmt"
)

func producer(itemCh chan<- int) {
	for i := 0; i < 10; i++ {
		itemCh <- i
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
		fmt.Println("Consumiu: ", n)
	}
}

func main() {
	itemCh := make(chan int)
	join := make(chan int)

	go producer(itemCh)
	go consume(itemCh, join)

	<-join
}
