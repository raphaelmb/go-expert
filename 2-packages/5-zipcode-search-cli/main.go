package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
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
	for _, zip := range os.Args[1:] {
		req, err := http.Get("http://viacep.com.br/ws/" + zip + "/json/")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error during request: %v\n", err)
		}
		defer req.Body.Close()
		res, err := io.ReadAll(req.Body)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error reading response: %v\n", err)
		}

		var data Zipcode
		err = json.Unmarshal(res, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing the response: %v\n", err)
		}
		file, err := os.Create("city.txt")
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
		}
		defer file.Close()

		_, err = file.WriteString(fmt.Sprintf("CEP: %s, Localidade: %s, UF: %s", data.Cep, data.Localidade, data.Uf))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
		}
		fmt.Println("File created successfully")
		fmt.Println(data.Localidade)
	}
}
