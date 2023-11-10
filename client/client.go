package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	response, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		fmt.Printf("Error while getting response from port -> %v\n", err)
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error while reading -> %v\n", err)
		return
	}
	fmt.Printf("Received %s\n", body)
}
