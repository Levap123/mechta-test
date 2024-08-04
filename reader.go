package main

import (
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
)

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func readData(path string, chanBuffer int) (<-chan Data, error) {
	dataCh := make(chan Data, chanBuffer)

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("can not open file: %v", err)
	}

	go func() {
		defer close(dataCh)
		defer file.Close()

		iter := jsoniter.Parse(jsoniter.ConfigDefault, file, 1024)

		for iter.ReadArray() {
			var data Data
			iter.ReadVal(&data)
			dataCh <- data
		}
	}()

	return dataCh, nil
}
