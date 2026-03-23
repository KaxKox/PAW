package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Current struct {
		Temp float64 `json:"temperature_2m"`
	} `json:"current"`
	Hourly struct {
		Time []string  `json:"time"`
		Temp []float64 `json:"temperature_2m"`
	} `json:"hourly"`
}

func main() {
	url := "https://api.open-meteo.com/v1/forecast?latitude=52.2297&longitude=21.0122&current=temperature_2m&hourly=temperature_2m"

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	var data Response
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		panic(err)
	}

	fmt.Printf("Aktualna temperatura: %.1f°C\n", data.Current.Temp)
	fmt.Println("\nPrognoza godzinowa:")

	for i := 0; i < 24; i++ {
		fmt.Printf("%s: %.1f°C\n", data.Hourly.Time[i][11:], data.Hourly.Temp[i])
	}
}
