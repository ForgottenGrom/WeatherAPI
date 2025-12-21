package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/redis/go-redis/v9"
)

var ctx = context.Background()
var rdb *redis.Client

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
type Geo struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

func main() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	if _, err := rdb.Ping(ctx).Result(); err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
	log.Println("Connected to Redis")

	http.HandleFunc("/", getWeather)
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	log.Fatalln(http.ListenAndServe(port, nil))

}
func getWeather(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	m := map[string]Geo{
		"Kyiv":    {Lat: 50.45, Lon: 30.52},
		"London":  {Lat: 51.51, Lon: -0.13},
		"NewYork": {Lat: 40.71, Lon: -74.01},
	}
	city := r.URL.Query().Get("city")
	val, err := rdb.Get(ctx, city).Result()
	if err == nil {
		log.Printf("Getting weather for %s from Redis", val)
		w.Write([]byte(val))
		return
	}
	geo, ok := m[city]
	if !ok {
		http.Error(w, "City not found", http.StatusNotFound)
		return
	}
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&current_weather=true", geo.Lat, geo.Lon)
	resp, err := http.Get(url)
	if err != nil {
		http.Error(w, "Failed to get weather data", http.StatusInternalServerError)
		return
	}
	var response OpenMeteoResponse
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read weather data", http.StatusInternalServerError)
		return
	}
	json.Unmarshal(body, &response)
	var desc string
	switch response.CurrentWeather.WeatherCode {
	case 0:
		desc = "Clear sky"
	default:
		desc = "Cloudy/Rainy"
	}
	weatherResult := WeatherData{
		City:        city,
		Temperature: response.CurrentWeather.Temperature,
		Description: desc,
	}

	d, _ := json.Marshal(weatherResult)
	err = rdb.Set(ctx, city, string(d), 30*time.Minute).Err()
	if err != nil {
		log.Printf("Failed to set weather data: %v", err)
	}
	w.Write(d)

}
