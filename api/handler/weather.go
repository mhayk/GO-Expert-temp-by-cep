package handler

import (
	"encoding/json"
	"net/http"

	"github.com/mhayk/GO-Expert-temp-by-cep/integration/address"
	"github.com/mhayk/GO-Expert-temp-by-cep/integration/weather"
	"gopkg.in/Nhanderu/brdoc.v1"
)

func getWeather() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		zipcode := r.URL.Query().Get("zipcode")

		// Validate Zipcode
		if brdoc.IsCEP(zipcode) == false {
			http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
			return
		}

		// Get address information
		zipcodeIntegration := address.ZipcodeIntegration{}
		addr, err := zipcodeIntegration.GetZipcode(zipcode)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if (address.Address{}) == *addr {
			http.Error(w, "can not find zipcode", http.StatusNotFound)
			return
		}

		// Get temperature information
		weatherIntegration := weather.WeatherIntegration{}
		temp, err := weatherIntegration.GetWeather(addr.City)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return the temperature in Celsius, Fahrenheit and Kelvin
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(temp)
	})
}

func NewWeatherHandler(r *http.ServeMux) {
	r.Handle("/temp-by-cep", getWeather())
}
