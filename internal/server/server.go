package server

import (
	"cachprax/internal/cache"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

type Config struct {
	Port   int
	Origin *url.URL
	Cache  *cache.Cache
}

func (cfg *Config) StartServer() error {
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port),
		Handler: mux,
	}

	mux.HandleFunc("/", cfg.proxyHandler)

	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("error starting the server: %v", err)
	}

	return nil
}

func (cfg *Config) proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Build the full URL for the origin server
	originURL := cfg.Origin
	originURL.Path = r.URL.Path
	originURL.RawQuery = r.URL.RawQuery

	// Check if the request is cached
	cacheKey := r.URL.String()
	if cfg.Cache.IsCached(cacheKey) {
		w.Header().Set("X-Cache", "HIT")
		cachedData := cfg.Cache.GetCached(cacheKey)
		if cachedData == nil {
			http.Error(w, "Cache entry is nil", http.StatusInternalServerError)
			return
		}
		_, err := w.Write(cachedData)
		if err != nil {
			http.Error(w, "Failed to write cached response", http.StatusInternalServerError)
			return
		}
		return
	}

	// Forward the request to the origin server
	client := &http.Client{}
	req, err := http.NewRequest(r.Method, originURL.String(), r.Body)
	if err != nil {
		http.Error(w, "Failed to create request to origin server", http.StatusInternalServerError)
		return
	}

	// Copy headers from the incoming request to the outgoing request
	for name, values := range r.Header {
		for _, value := range values {
			req.Header.Add(name, value)
		}
	}

	// Send the request to the origin server
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, "Failed to reach the origin server", http.StatusBadGateway)
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	// Set the response headers and status code
	w.Header().Set("X-Cache", "MISS")
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}
	w.WriteHeader(resp.StatusCode)

	// Read the response body into a buffer
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}

	// Cache the response body
	cfg.Cache.SetCached(cacheKey, bodyBytes)

	// Write the response body to the client
	_, err = w.Write(bodyBytes)
	if err != nil {
		http.Error(w, "Failed to write response body", http.StatusInternalServerError)
		return
	}
}
