package main

import "fmt"

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}

func main() {
	weather := WeatherData{"London", 12.5, "Fog"}
	fmt.Println(weather)
}
