package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"
)

func writeJSON(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(map[string]interface{}{
		"result": data,
	})
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func writeError(w http.ResponseWriter, msg string, code int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if err := json.NewEncoder(w).Encode(map[string]string{"error": msg}); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
	}
}

func checkMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return false
	}
	return true
}

func queryParseUserID(r *http.Request) (uint, error) {
	query := r.URL.Query().Get("user_id")
	userID, err := strconv.Atoi(query)
	if err != nil {
		return 0, err
	}
	return uint(userID), nil
}

func queryParseDate(r *http.Request) (time.Time, error) {
	query := r.URL.Query().Get("date")
	t, err := time.Parse("2006-01-02", query)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}
