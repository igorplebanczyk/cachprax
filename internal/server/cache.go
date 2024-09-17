package server

import "net/http"

func (cfg *Config) clearCacheHandler(w http.ResponseWriter, r *http.Request) {
	cfg.Cache.Clear()
	w.WriteHeader(http.StatusNoContent)
}
