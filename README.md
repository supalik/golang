# golang
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sync"
)

var msg = []string{
	"test1",
	"test2",
	"test3",
	"test4",
	"test5",
	"test6",
	"test7",
}

func producer(wg *sync.WaitGroup, data chan<- string, i int) {
	defer wg.Done()
	for _, val := range msg {
		data <- val
		fmt.Printf("val %v producer %v \n", val, i)
	}
	//close(data)
}
func consumer(i int, data <-chan string, wc *sync.WaitGroup) {
	defer wc.Done()
	for v := range data {
		fmt.Printf("consumed %v by  %v \n", v, i)
	}
	//status <- true
}

func main() {
	fmt.Println("Hello, producer and consumer")
	wg := sync.WaitGroup{}
	wc := sync.WaitGroup{}
	data := make(chan string)
	//status := make(chan bool)
	for j := 1; j < 5; j++ {
		wg.Add(1)
		go producer(&wg, data, j)

	}

	for i := 1; i < 6; i++ {
		go consumer(i, data, &wc)
	}
	wg.Wait()
	//<-status
	close(data)
	wc.Wait()
}
