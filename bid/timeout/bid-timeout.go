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

func bid(item int) chan Bid {
	bid := make(chan Bid)
	go func() {
		time.Sleep(time.Second * 5)
		bid <- Bid{item, item + 1, false}
	}()
	return bid
}

func handle(nServers int, itemCh <-chan int, timeoutSecs int) chan Bid {
	bidCh := make(chan Bid)
	join := make(chan int)
	tick := time.Tick(time.Duration(timeoutSecs) * time.Second)

	for i := 0; i < nServers; i++ {
		go func() {
			for item := range itemCh {
				select {
				case <-tick:
					b := Bid{item, -1, true}
					bidCh <- b
				case b := <-bid(item):
					bidCh <- b
				}
			}
			join <- 1
		}()
	}
	go func() {
		for i := 0; i < nServers; i++ {
			<-join
		}
		close(join)
		close(bidCh)
	}()

	return bidCh
}

func generateInput(itemCh chan<- int) {
	for i := 0; i < 15; i++ {
		itemCh <- i
	}
	close(itemCh)
}

func main() {
	itemCh := make(chan int)
	go generateInput(itemCh)

	bidCh := handle(5, itemCh, 2)
	for bid := range bidCh {
		fmt.Println("bid: ", bid)
	}

	fmt.Print("Done")
}
