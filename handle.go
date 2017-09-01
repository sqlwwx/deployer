package main

import (
	"log"
	"net/http"
	"strings"
)

func handle(w http.ResponseWriter, req *http.Request) {
	if origin := req.Header.Get("Origin"); origin != "" {
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Origin", origin)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers",
			"Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	}

	// Stop here for a Preflighted OPTIONS request.
	if req.Method == "OPTIONS" {
		return
	}
	req.ParseForm()
	secret := req.FormValue("secret")

	pathInfo := strings.Trim(req.URL.Path, "/")
	parts := strings.Split(pathInfo, "/")
	if len(parts) > 0 {
		id := parts[0]
		for _, item := range cfg.Items {
			if id == item.Id && item.Secret == secret {
				log.Println(id)
				err := runScript(&item)
				if err != nil {
					log.Printf("run script error: %s\n", err)
				} else {
					w.Write([]byte("success"))
				}
				break
			}
		}
	} else {
		w.Write([]byte("require id"))
	}
}
