package main

import (
	"cachprax/cmd"
	"cachprax/internal/cache"
	"fmt"
	"net/http"
	"net/url"
	"os"
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

	cliApp := cmd.NewApp()
	err := cliApp.Run(os.Args)
	if err != nil {
		fmt.Printf("error running cli: %v", err)
	}

	err = server.ListenAndServe()
	if err != nil {
		return
	}
}
