package main

import (
	"cachprax/internal/cache"
	"fmt"
	"net/http"
	"net/url"
	"time"
)

type Config struct {
	Port   int
	Origin url.URL
	Cache  *cache.Cache
}

func main() {
	cfg := Config{
		Port: 8080,
		Origin: url.URL{
			Scheme: "http",
			Host:   "httpbin.org",
		},
		Cache: cache.NewCache(5*time.Minute, 10*time.Minute),
	}

	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	mux.HandleFunc("/", cfg.proxyHandler)

	err := server.ListenAndServe()
	if err != nil {
		return
	}
}
