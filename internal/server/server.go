package server

import (
	"cachprax/internal/cache"
	"fmt"
	"net/http"
	"net/url"
)

type Config struct {
	Port   int
	Origin *url.URL
	Cache  *cache.Cache
}

func (cfg *Config) StartServer() error {
	// server for cache operations
	cacheMux := http.NewServeMux()
	cacheServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", 3001),
		Handler: cacheMux,
	}

	cacheMux.HandleFunc("/cache/clear", cfg.clearCacheHandler)

	go func() {
		err := cacheServer.ListenAndServe()
		if err != nil {
			fmt.Printf("error starting the cache server: %v\n", err)
		}
	}()

	// server for proxy operations
	proxyMux := http.NewServeMux()
	proxyServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: proxyMux,
	}

	proxyMux.HandleFunc("/", cfg.proxyHandler)

	err := proxyServer.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error starting the server: %v", err)
	}

	return nil
}
