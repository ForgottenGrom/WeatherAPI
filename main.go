package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func main() {
	weather := WeatherData{"London", 12.5, "Fog"}

	jsonData, err := json.Marshal(weather)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
}
