package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

const (
	BrasilAPIURL = "https://brasilapi.com.br/api/cep/v1/"
	ViaCEPURL    = "http://viacep.com.br/ws/"
)

func main() {
	address := make(chan string, 2)

	var cep string
	fmt.Println("Enter zip code for address lookup:")
	fmt.Scan(&cep)

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	go func() {
		RequestBrasilApi(cep, address, ctx)
		cancel()
	}()

	go func() {
		RequestViaCep(cep, address, ctx)
		cancel()
	}()

	for {
		select {
		case msg := <-address:
			fmt.Printf(msg + "\n")
			return

		case <-time.After(1 * time.Second):
			println("timeout")
			return
		}
	}
}

func RequestBrasilApi(cep string, addressChan chan<- string, ctx context.Context) {
	url := BrasilAPIURL + cep

	resp, err := fetchAddress(url, ctx)
	if err != nil {
		addressChan <- fmt.Sprintf("Error fetching address from Brasil API: %v", err)
		return
	}

	addressChan <- fmt.Sprintf("Address from Brasil API: %s", resp)
}

func RequestViaCep(cep string, addressChan chan<- string, ctx context.Context) {
	url := ViaCEPURL + cep + "/json/"

	resp, err := fetchAddress(url, ctx)
	if err != nil {
		addressChan <- fmt.Sprintf("Error fetching address from ViaCEP: %v", err)
		return
	}

	addressChan <- fmt.Sprintf("Address from ViaCEP: %s", resp)
}

func fetchAddress(url string, ctx context.Context) (string, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return "Could not make new request with context", err
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		return "Error fetching address", err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return "Error reading response body:", err
	}
	return string(body), nil
}
