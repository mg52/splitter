package splitter

import (
	"fmt"
	"testing"
)

type Temperature struct {
	Id    int     `json:"id"`
	Value float64 `json:"value"`
	Type  string  `json:"type"`
}

// TestWhere tests splitter.TestWhere
func TestWhere(t *testing.T) {
	var temperatures []Temperature

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

	clauses := []Clause{
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
	actual, err := Where(temperatures, clauses)
	if err != nil {
		fmt.Println(err.Error())
	}

	var expected []Temperature
	expected = append(expected, Temperature{
		Id:    3,
		Value: 94.9,
		Type:  "Fahrenheit",
	})
	expected = append(expected, Temperature{
		Id:    4,
		Value: 80.2,
		Type:  "Fahrenheit",
	})

	if len(expected) != len(actual) {
		t.Errorf("expected and actual lengths are not equal. actual length: %v, expected length: %v", len(actual), len(expected))

	}
	for i := 0; i < len(expected); i++ {
		if expected[i].Id != actual[i].Id ||
			expected[i].Value != actual[i].Value ||
			expected[i].Type != actual[i].Type {
			t.Errorf("actual %v, expected %v", actual[i], expected[i])
		}
	}
}
