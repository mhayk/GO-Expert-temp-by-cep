package address

import (
	"encoding/json"
	"net/http"
)

type Address struct {
	CEP          string `json:"cep"`
	Street       string `json:"logradouro"`
	Complement   string `json:"complemento"`
	Neighborhood string `json:"bairro"`
	City         string `json:"localidade"`
	State        string `json:"uf"`
	IBGE         string `json:"ibge"`
	GIA          string `json:"gia"`
	DDD          string `json:"ddd"`
	SIAFI        string `json:"siafi"`
}

type ZipcodeInterface interface {
	GetZipcode(zipcode string) (*Address, error)
}

// ZipcodeInterface is the interface that wraps the basic methods for a zipcode service.
type ZipcodeIntegration struct{}

// GetZipcode is a method that returns the address information for a given zipcode.
func (z *ZipcodeIntegration) GetZipcode(zipcode string) (*Address, error) {
	req, err := http.NewRequest("GET", "https://viacep.com.br/ws/"+zipcode+"/json/", nil)

	if err != nil {
		return nil, err
	}

	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)

	if err != nil {
		return nil, err
	}

	return &address, nil
}
