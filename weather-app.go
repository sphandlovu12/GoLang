package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const (
	// Replace with your OpenWeatherMap API key
	apiKey = "YOUR_API_KEY"
	city   = "New York" // Replace with the city of your choice
)

type WeatherData struct {
	Main struct {
		Temp float64 `json:"temp"`
	} `json:"main"`
	Weather []struct {
		Description string `json:"description"`
	} `json:"weather"`
}

func getWeather() (*WeatherData, error) {
	url := fmt.Sprintf("http://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s", city, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var weatherData WeatherData
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, err
	}

	return &weatherData, nil
}

func main() {
	ticker := time.NewTicker(1 * time.Hour) // Update every hour
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			weatherData, err := getWeather()
			if err != nil {
				fmt.Printf("Error: %v\n", err)
			} else {
				description := weatherData.Weather[0].Description
				temperature := weatherData.Main.Temp
				fmt.Printf("Weather in %s: %s, Temperature: %.2f°C\n", city, description, temperature-273.15)
			}
		}
	}
}
