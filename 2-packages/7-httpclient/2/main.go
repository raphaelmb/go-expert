package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
)

func main() {
	c := http.Client{}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "john"}`))
	res, err := c.Post("https://www.google.com", "application/json", jsonVar)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)

}
