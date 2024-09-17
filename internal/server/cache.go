package server

import (
	"net/http"
	"strconv"
)

func (cfg *Config) clearCacheHandler(w http.ResponseWriter, r *http.Request) {
	cfg.Cache.Clear()
	w.WriteHeader(http.StatusNoContent)
}

func (cfg *Config) countCacheHandler(w http.ResponseWriter, r *http.Request) {
	count := cfg.Cache.Count()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte(strconv.Itoa(count)))
	if err != nil {
		return
	}
}
