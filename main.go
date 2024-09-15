package main

import (
	"net/http"
	"net/url"
	"strconv"
)

type Config struct {
	Port   int
	Origin url.URL
}

func main() {
	cfg := Config{
		Port: 8080,
		Origin: url.URL{
			Scheme: "http",
			Host:   "httpbin.org",
		},
	}
	port := strconv.Itoa(cfg.Port)

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/", cfg.proxyHandler)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
