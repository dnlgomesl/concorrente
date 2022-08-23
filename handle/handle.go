package main

// Ta dando deadlock e eu não sei o pq

import (
	"fmt"
	"math/rand"
	"time"
)

type LightSwitch struct {
	count int
	mutex chan int
}

func LightSwitch_init() LightSwitch {
	return LightSwitch{count: 0, mutex: make(chan int, 1)}
}

func (ls LightSwitch) lock(s chan int) {
	ls.mutex <- 0
	ls.count += 1
	if ls.count == 1 {
		s <- 0
	}
	<-ls.mutex
}

func (ls LightSwitch) unlock(s chan int) {
	ls.mutex <- 0
	ls.count -= 1
	if ls.count == 0 {
		<-s
	}
	<-ls.mutex
}

func generateInput(inputs chan int) {
	rand.Seed(26)
	for i := 0; i < 6; i++ {
		n := rand.Intn(2)
		inputs <- n
	}
	close(inputs)
}

func handle(zeros LightSwitch, uns LightSwitch, turnstile chan int, s0 chan int, s1 chan int, sem_lock chan int, req int) {
	if req == 0 {
		<-turnstile
		zeros.lock(sem_lock)
		turnstile <- 1
		<-s0
		time.Sleep(time.Duration(rand.Intn(5)) * time.Second)
		fmt.Println("Req type: ", req)
		s0 <- 1
		zeros.unlock(sem_lock)
	} else {
		<-turnstile
		uns.lock(sem_lock)
		turnstile <- 1
		<-s1
		fmt.Println("Req type: ", req)
		s1 <- 1
		uns.unlock(sem_lock)
	}
}

func consume(zeros LightSwitch, uns LightSwitch, turnstile chan int, s0 chan int, s1 chan int, sem_lock chan int, reqs chan int, join chan int) {
	for req := range reqs {
		handle(zeros, uns, turnstile, s0, s1, sem_lock, req)
	}
	join <- 1
	close(join)
}

func main() {
	n := 2
	sem_lock := make(chan int, 1)
	s0 := make(chan int, n)
	s1 := make(chan int, n)
	turnstile := make(chan int, 1)
	zeros := LightSwitch_init()
	uns := LightSwitch_init()

	reqs := make(chan int)
	join := make(chan int)

	go generateInput(reqs)
	go consume(zeros, uns, turnstile, s0, s1, sem_lock, reqs, join)

	<-join
}
