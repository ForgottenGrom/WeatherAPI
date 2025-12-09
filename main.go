package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type WeatherData struct {
	City        string  `json:"city"`
	Temperature float64 `json:"temperature"`
	Description string  `json:"description"`
}
type OpenMeteoResponse struct {
	CurrentWeather struct {
		Temperature float64 `json:"temperature"`
		WeatherCode float64 `json:"weathercode"`
	} `json:"current_weather"`
}

func main() {
	weather := WeatherData{"London", 12.5, "Fog"}

	jsonData, err := json.Marshal(weather)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(jsonData))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		resp, _ := http.Get("https://api.open-meteo.com/v1/forecast?latitude=50.45&longitude=30.52&current_weather=true")
		var response OpenMeteoResponse
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		json.Unmarshal(body, &response)
		var desc string
		switch response.CurrentWeather.WeatherCode {
		case 0:
			desc = "Clear sky"
		default:
			desc = "Cloudy/Rainy"
		}
		weatherResult := WeatherData{
			City:        "Kyiv",
			Temperature: response.CurrentWeather.Temperature,
			Description: desc,
		}
		d, _ := json.Marshal(weatherResult)
		w.Write(d)
	})

	http.ListenAndServe(":8080", nil)

}
