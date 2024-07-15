package main

import (
	"fmt"
	"sync"

	"github.com/luizfelipe94/worker-poc/data"
	"github.com/luizfelipe94/worker-poc/send"
)

func main() {
	var wg sync.WaitGroup
	totalWorkers := 50

	fmt.Println("process starting...")

	channel := data.Generator(100)
	for i := 1; i <= totalWorkers; i++ {
		wg.Add(1)
		go worker(i, channel, &wg)
	}

	wg.Wait()
	fmt.Println("stopping process...")
}

func worker(id int, channel <-chan data.Record, wg *sync.WaitGroup) {
	defer wg.Done()
	for message := range channel {
		send.Send(id, message)
	}
}
