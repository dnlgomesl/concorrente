package main

import (
	"fmt"
	"time"
)

type Bid struct {
	item      int
	bidValue  int
	bidFailed bool
}

func bid(item int) Bid {
	time.Sleep(time.Second * 5)
	return Bid{item, item + 1, false}
}

func handle(nServers int, itemCh <-chan int) chan Bid {
	bidCh := make(chan Bid)
	join := make(chan int)

	for i := 0; i < nServers; i++ {
		go func() {
			for item := range itemCh {
				bidCh <- bid(item)
			}
			join <- 1
		}()
	}
	go func () {
		for i :=0; i < nServers; i++{
			<- join
		}
		close(join)
		close(bidCh)
	} ()

	return bidCh
}

func generateInput(itemCh chan<- int) {
	for i := 0; i < 15; i++{
		itemCh <- i
	}
	close(itemCh)
}

func main() {
	itemCh := make(chan int)
	go generateInput(itemCh)

	bidCh := handle(5, itemCh)
	for bid := range bidCh {
		fmt.Println("bid: ", bid)
	}

	fmt.Print("Done")
}
