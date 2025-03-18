package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func GetCep(apiName string, url string, ch chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Panic(err.Error())
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panic(err.Error())
	}
	ch <- fmt.Sprintf("Retorno da %s: %s", apiName, string(body))
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go GetCep("brasilapi", "https://brasilapi.com.br/api/cep/v1/01153000", ch1)
	go GetCep("viacep", "https://viacep.com.br/ws/01153000/json/", ch2)

	select {
	case response := <-ch1:
		fmt.Println(response)
	case response := <-ch2:
		fmt.Println(response)
	case <-time.After(time.Second):
		fmt.Println("timeout")
	}
}
