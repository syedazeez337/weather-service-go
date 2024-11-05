package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	configs "weatherservice/config"
)

func GetWeather(city string) (*WeatherData, error) {
	apiKey, err := configs.GetAPIKey()
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.openweathermap.org/data/2.5/weather?q=%s&appid=%s&units=metric", city, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %v", resp.Status)
	}

	var data WeatherData
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return &data, nil
}

func DisplayWeather(data *WeatherData) {
	// fmt.Println(data)
	fmt.Printf("Current weather in %s\n", data.Name)
	fmt.Printf("Temperature: %.2fC\n", data.Main.Temp)
	fmt.Printf("Condition: %s\n", data.Weather[0].Description)
}
