# splitter

Go LINQ Where

## Installation

	go get github.com/mg52/splitter@v1.0.1

## Usage: 
```go
package main

import (
	"fmt"
	"github.com/mg52/splitter"
)

type Temperature struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

var temperatures []Temperature

func main() {
	temperatures = append(temperatures, Temperature{
		Id:    0,
		Value: 15.3,
		Type:  "Celsius",
	})
	temperatures = append(temperatures, Temperature{
		Id:    1,
		Value: 27.4,
		Type:  "Celsius",
	})
	temperatures = append(temperatures, Temperature{
		Id:    2,
		Value: 16.5,
		Type:  "Celsius",
	})
	temperatures = append(temperatures, Temperature{
		Id:    3,
		Value: 94.9,
		Type:  "Fahrenheit",
	})
	temperatures = append(temperatures, Temperature{
		Id:    4,
		Value: 80.2,
		Type:  "Fahrenheit",
	})
	temperatures = append(temperatures, Temperature{
		Id:    5,
		Value: 75.4,
		Type:  "Fahrenheit",
	})

	clauses := []splitter.Clause{
		{
			Key:    "Type",
			Method: "==",
			Val:    "Fahrenheit",
		},
		{
			Key:    "Value",
			Method: ">=",
			Val:    75.6,
		},
	}

	actual, err := splitter.Where(temperatures, clauses)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(actual)
}
```
