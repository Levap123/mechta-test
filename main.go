package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("usage: go run . <filename> <number_of_workers>")
		return
	}

	filename := os.Args[1]
	workersNum, err := strconv.Atoi(os.Args[2])
	if err != nil || workersNum <= 0 {
		log.Fatalf("provide valid workers number: %v", err)
	}

	results := make(chan int, workersNum)

	dataCh, err := readData(filename, workersNum)
	if err != nil {
		log.Fatalf("read data: %v", err)
	}

	var wg sync.WaitGroup
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(dataCh, results, &wg)
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	totalSum := 0
	for sum := range results {
		totalSum += sum
	}

	fmt.Println("Total sum:", totalSum)
}
