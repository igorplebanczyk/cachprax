package main

import (
	"io"
	"net/http"
)

func (cfg *Config) proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Build the full URL for the origin server
	originURL := cfg.Origin
	originURL.Path = r.URL.Path
	originURL.RawQuery = r.URL.RawQuery

	// Check if the request is cached
	cacheKey := r.URL.String()
	if cfg.Cache.IsCached(cacheKey) {
		w.Header().Set("X-Cache", "HIT")
		_, err := w.Write(cfg.Cache.GetCached(cacheKey))
		if err != nil {
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

	// Copy the headers from the incoming request to the outgoing request
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
			http.Error(w, "Failed to close response body", http.StatusInternalServerError)
		}
	}(resp.Body)

	w.Header().Set("X-Cache", "MISS")

	// Copy the headers from the origin server’s response
	for name, values := range resp.Header {
		for _, value := range values {
			w.Header().Add(name, value)
		}
	}

	// Set the response status code
	w.WriteHeader(resp.StatusCode)

	// Cache the response body
	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		http.Error(w, "Failed to read response body", http.StatusInternalServerError)
		return
	}
	cfg.Cache.SetCached(cacheKey, bodyBytes)

	// Copy the response body from the origin server to the client
	_, err = io.Copy(w, resp.Body)
	if err != nil {
		http.Error(w, "Failed to copy response body", http.StatusInternalServerError)
		return
	}
}
