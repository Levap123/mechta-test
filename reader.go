package main

import (
	"fmt"
	"os"

	jsoniter "github.com/json-iterator/go"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type Data struct {
	A int `json:"a"`
	B int `json:"b"`
}

func readData(path string) ([]Data, error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("os.ReadFile: %w", err)
	}

	var data []Data

	if err := json.Unmarshal(bytes, &data); err != nil {
		return nil, fmt.Errorf("json.Unmarshal: %w", err)
	}

	return data, nil
}
