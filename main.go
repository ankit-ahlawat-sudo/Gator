package main

import (
	"log"
	"fmt"

	"github.com/ankit-ahlawat-sudo/Gator/internal/config"
)

func main() {
	cfg, err:= config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	if err := cfg.SetUser("AnkitAhlawat"); err != nil {
		log.Fatalf("error writing config: %v", err)
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}
	fmt.Printf("%+v\n", cfg)
	
}
