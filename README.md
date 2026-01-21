# Weather API Service

This is a weather API project that fetches and returns weather data from a 3rd party API. This project is based on the [Weather API Project Idea](https://roadmap.sh/projects/weather-api-wrapper-service) from the [roadmap.sh](https://roadmap.sh/dashboard) platform.

## Project Overview

Instead of relying on our own weather data, this project aims to build a weather API that fetches and returns weather data from a 3rd party API. This project demonstrates how to work with external APIs, implement server-side caching, and manage environment variables in Go.

The project uses the **Open-Meteo API** as the 3rd party provider to fetch real-time weather data. This API is free, requires no API key for basic usage, and provides accurate forecast data.

For caching, the project uses **Redis**. To optimize performance and reduce external API calls, the weather data for a specific city is cached for **30 minutes**. If a user requests data for a city that is already in the cache, the service returns the data instantly from memory.

## Technologies Used

* **Programming Language:** Go (Golang)
* **Standard Library:** `net/http` (No heavy frameworks used)
* **Caching:** Redis (`go-redis/v9`)
* **External API:** [Open-Meteo](https://open-meteo.com/)

## Features

* **3rd Party Integration:** Fetches real-time weather data from Open-Meteo.
* **Redis Caching:** Caches responses for 30 minutes to improve response times and save bandwidth.
* **Rate Limiting:** Implements a fixed-window rate limiter (5 requests per minute per IP) to prevent abuse.
* **Environment Variables:** Configurable `PORT` and `REDIS_ADDR` for flexible deployment.
* **Error Handling:** Robust handling of timeout errors, JSON parsing errors, and external API failures.

## Getting Started

### Prerequisites
* [Go](https://go.dev/dl/) (version 1.19 or higher)
* [Redis](https://redis.io/) server running locally or remotely

### Installation & Run

1.  **Clone the repository:**
    ```bash
    git clone [https://github.com/your-username/go-weather-api.git](https://github.com/your-username/go-weather-api.git)
    cd go-weather-api
    ```

2.  **Install dependencies:**
    ```bash
    go mod tidy
    ```

3.  **Start Redis:**
    Ensure your Redis server is running on the default port (or configure via env vars).
    ```bash
    redis-server
    ```

4.  **Run the application:**
    ```bash
    go run main.go
    ```

The Weather API will be available at `http://localhost:8080`.

## Usage Example

To get the weather for a specific city, send a GET request:

```http
GET http://localhost:8080/?city=Kyiv
```
Response:

```JSON
{
  "city": "Kyiv",
  "temperature": 18.2,
  "description": "Cloudy "
}
```
## Contribution
This project is part of the [roadmap.sh](https://roadmap.sh/dashboard) platform, which is designed to help developers learn and grow their skills. If you find any issues or have suggestions for improvement, please feel free to contribute to the project by submitting a pull request or opening an issue.
## License
This project is licensed under the MIT License.
