package integration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mhayk/GO-Expert-temp-by-cep/api/handler"
	"github.com/h2non/gock"
	"github.com/stretchr/testify/assert"
)

func TestShouldReturnTemperatureInCelsiusFahrenheitAndKelvin(t *testing.T) {
	defer gock.Off()

	gock.New("https://viacep.com.br").
		Get("/ws/69304350/json/").
		Reply(200).
		JSON(map[string]string{
			"cep":         "69304-350",
			"logradouro":  "Avenida Mário Homem de Melo",
			"complemento": "de 729/730 a 2387/2388",
			"bairro":      "Mecejana",
			"localidade":  "Boa Vista",
			"uf":          "RR",
			"ibge":        "1400100",
			"gia":         "",
			"ddd":         "95",
			"siafi":       "0301",
		})

	gock.New("https://api.weatherapi.com").
		Get("/v1/current.json").
		Reply(200).
		JSON(map[string]interface{}{
			"location": map[string]interface{}{
				"name": "Boa Vista",
			},
			"current": map[string]interface{}{
				"temp_c": 20.0,
				"temp_f": 68.0,
			},
		})

	r := http.NewServeMux()

	handler.NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/temp-by-cep?zipcode=69304350", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 200 {
		t.Errorf("Expected status code 200, got %d", w.Code)
	}

	expected := `{"temp_C":20,"temp_F":68,"temp_K":293}`

	assert.Equal(t, expected, strings.TrimRight(w.Body.String(), "\n"))
}

func TestShouldReturnErrorWhenZipcodeIsInvalid(t *testing.T) {
	r := http.NewServeMux()

	handler.NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/temp-by-cep?zipcode=123", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 422 {
		t.Errorf("Expected status code 422, got %d", w.Code)
	}

	expected := "invalid zipcode\n"

	assert.Equal(t, expected, w.Body.String())
}

func TestShouldReturnErrorWhenZipcodeIsNotFound(t *testing.T) {
	defer gock.Off()

	gock.New("https://viacep.com.br").
		Get("/ws/12345678/json/").
		Reply(404).
		JSON(map[string]bool{
			"erro": true,
		})

	r := http.NewServeMux()

	handler.NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/temp-by-cep?zipcode=12345678", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 404 {
		t.Errorf("Expected status code 404, got %d", w.Code)
	}

	expected := "can not find zipcode\n"

	assert.Equal(t, expected, w.Body.String())
}

func TestShouldReturnErrorWhenWeatherAPIIsUnavailable(t *testing.T) {
	defer gock.Off()

	gock.New("https://viacep.com.br").
		Get("/ws/69304350/json/").
		Reply(200).
		JSON(map[string]string{
			"cep":         "69304-350",
			"logradouro":  "Avenida Mário Homem de Melo",
			"complemento": "de 729/730 a 2387/2388",
			"bairro":      "Mecejana",
			"localidade":  "Boa Vista",
			"uf":          "SP",
			"ibge":        "1400100",
			"gia":         "",
			"ddd":         "95",
			"siafi":       "0301",
		})

	gock.New("https://api.weatherapi.com").
		Get("/v1/current.json").
		Reply(500)

	r := http.NewServeMux()

	handler.NewWeatherHandler(r)

	req := httptest.NewRequest("GET", "/temp-by-cep?zipcode=69304350", nil)

	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	if w.Code != 500 {
		t.Errorf("Expected status code 500, got %d", w.Code)
	}
}
