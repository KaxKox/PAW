package main

import (
	"fmt"
	"log"
	
	"config/config"
)

func main() {
	cfg, err := config.Load("config.yaml")
	if err != nil {
		log.Fatalf("Błąd krytyczny: %v\n", err)
	}

	fmt.Printf("Serwer: %s:%d\n", cfg.Server.Host, cfg.Server.Port)
	
	fmt.Printf("Baza danych:\nuser: %s\nhasło: %s\n", cfg.Database.User, cfg.Database.Pass)
}