package weather

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
)

type CurrentWeather struct {
	TemperatureC float64 `json:"temp_c"`
	TemperatureF float64 `json:"temp_f"`
}

type WeatherData struct {
	Current CurrentWeather `json:"current"`
}

type Weather struct {
	Temp_C float64 `json:"temp_C"`
	Temp_F float64 `json:"temp_F"`
	Temp_K float64 `json:"temp_K"`
}

type WeatherIntegration struct{}

func (w *WeatherIntegration) GetWeather(city string) (*Weather, error) {
	req, err := http.NewRequest("GET", "https://api.weatherapi.com/v1/current.json?key="+os.Getenv("WEATHER_API_KEY")+"&q="+url.QueryEscape(city)+"&aqi=no", nil)
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	// Here we would parse the response from the weather API
	var weather WeatherData
	err = json.NewDecoder(resp.Body).Decode(&weather)

	if err != nil {
		return nil, err
	}

	// and return the temperature in Celsius, Fahrenheit and Kelvin
	return &Weather{
		Temp_C: weather.Current.TemperatureC,
		Temp_F: weather.Current.TemperatureF,
		Temp_K: weather.Current.TemperatureC + 273,
	}, nil
}
