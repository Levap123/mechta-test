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

	data, err := readData(filename)
	if err != nil {
		log.Fatalf("read data: %v", err)
	}

	dataChan := make(chan Data, workersNum)
	results := make(chan int, workersNum)

	var wg sync.WaitGroup
	for i := 0; i < workersNum; i++ {
		wg.Add(1)
		go worker(dataChan, results, &wg)
	}

	go func() {
		for _, d := range data {
			dataChan <- d
		}
		close(dataChan)
	}()

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
