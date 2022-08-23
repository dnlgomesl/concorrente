package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"time"
)

func read(filter chan string, join chan int) {
	for file := range filter {
		fmt.Println(file)
	}
	join <- 1
}

func find(root string, filter chan string, i int) {
	files, err := ioutil.ReadDir(root)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		new := root + "/" + file.Name()
		// sleep so pra ver a execução direitinho
		time.Sleep(5 * time.Second)
		if !file.IsDir() {
			filter <- new
		} else {
			find(new, filter, i+1)
		}
	}
	if i == 0 {
		close(filter)
	}
}

func main() {
	root := "./dir1/"

	filter := make(chan string)
	join := make(chan int)

	go find(root, filter, 0)
	go read(filter, join)

	<-join
	close(join)
}
