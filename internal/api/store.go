package api

import (
	"net/http"

	"github.com/semihsemih/kv-store/internal/storage"
)

// Set Handle request to set new key-value in storage
func Set(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			RespondWithError(w, http.StatusBadRequest, "Url Param 'key' is missing")
			return
		}

		values, ok := r.URL.Query()["value"]
		if !ok || len(values[0]) < 1 {
			RespondWithError(w, http.StatusBadRequest, "Url Param 'value' is missing")
			return
		}

		resp := s.Set(keys[0], values[0])

		RespondWithJSON(w, http.StatusOK, resp)
	}
}

// Get Handle request to get value from storage by given key
func Get(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		keys, ok := r.URL.Query()["key"]
		if !ok || len(keys[0]) < 1 {
			RespondWithError(w, http.StatusBadRequest, "Url Param 'key' is missing")
			return
		}

		resp, err := s.Get(keys[0])
		if err != nil {
			RespondWithError(w, http.StatusNotFound, err.Error())
			return
		}

		RespondWithJSON(w, http.StatusOK, resp)
	}
}

// Flush Handle request to flush all data in storage
func Flush(s storage.Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.Flush()

		RespondWithJSON(w, http.StatusOK, "This storage is clear.")
	}
}
