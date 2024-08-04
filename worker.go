package main

import "sync"

func worker(dataChan <-chan Data, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	sum := 0
	for data := range dataChan {
		sum += data.A + data.B
	}

	results <- sum
}
