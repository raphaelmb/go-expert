package main

import (
	"encoding/json"
	"io"
	"net/http"
)

type Zipcode struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	http.HandleFunc("/", SearchZipcodeHandler)
	http.ListenAndServe(":8080", nil)
}

func SearchZipcodeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	zipParam := r.URL.Query().Get("cep")
	if zipParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	zip, err := SearchZipcode(zipParam)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(zip)
}

func SearchZipcode(zip string) (*Zipcode, error) {
	req, err := http.Get("http://viacep.com.br/ws/" + zip + "/json/")
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	body, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}

	var z Zipcode
	err = json.Unmarshal(body, &z)
	if err != nil {
		return nil, err
	}

	return &z, nil
}
