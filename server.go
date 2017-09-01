package main

import (
	"log"
	"net/http"
	"os"
)

var cfg Config

func main() {

	if len(os.Args) < 2 {
		println("Usage: webhook <ConfigFile>\n")
		return
	}

	err := readConfig(os.Args[1], &cfg)
	if err != nil {
		log.Println("Unmarshal config failed:", err)
		return
	}

	http.HandleFunc("/", handle)
	log.Println("listen at", cfg.BindHost)
	log.Fatal(http.ListenAndServe(cfg.BindHost, nil))
}
