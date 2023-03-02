package main

import (
	"encoding/json"
	"fmt"
)

//json- simple data interchange format.
//marshal -> encode
//unmarshal -> decode

type Car struct {
	Model       string `json:"model"`
	Description string `json:"description"`
}
type Dimensions struct {
	Height int
	Width  int
}
type CarWithDimension struct {
	Model       string     `json:"model"`
	Description string     `json:"description"`
	Dimensions  Dimensions `json:"dimensions"`
}

func main() {
	//unmarshal row json data
	myJson := `{"model": "cat","description": "this is car description"}`
	var car Car
	json.Unmarshal([]byte(myJson), &car)
	fmt.Printf("Car: %+v", car)
	// nested objects
	myJsonWithDimesions := `{"model": "cat","description": "this is car description","dimensions":{"height":24,"width":100}}`
	var carWithDimensions CarWithDimension
	json.Unmarshal([]byte(myJsonWithDimesions), &carWithDimensions)
	fmt.Printf("carWithDimensions: %+v", carWithDimensions)

	//Decoding JSON to Maps - Unstructured Data
	randomJson := `{"car":{"model": "cat","description": "this is car description","color":"red"}}`
	var result map[string]any
	json.Unmarshal([]byte(randomJson), &result)

	// Marshaling Structured Data
	response := &Car{
		Model:       "2023",
		Description: "car description",
	}
	data, _ := json.Marshal(response)
	fmt.Println(string(data))
}
