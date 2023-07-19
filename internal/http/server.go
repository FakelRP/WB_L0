package http

import (
	"WB_L0/internal/models"
	"encoding/json"
	"net/http"
	"strconv"
)

// HTTP обработчик для получения данных из кеша по ID.
func getDataFromCacheHandler(cache *models.Cache) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "Invalid ID", http.StatusBadRequest)
			return
		}

		cache.RLock()
		message, ok := cache.Data[strconv.Itoa(id)]
		cache.RUnlock()

		if !ok {
			http.Error(w, "Message not found", http.StatusNotFound)
			return
		}

		jsonData, err := json.Marshal(message)
		if err != nil {
			http.Error(w, "Failed to marshal data", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonData)
	}
}
