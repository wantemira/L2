package controller

import (
	"calendar/pkg/models"
	"encoding/json"
	"net/http"
	"strconv"
)

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodPost) {
		return
	}

	var req models.EventRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, "bad request", http.StatusBadRequest)
		return
	}

	event, err := h.service.Create(&req)
	if err != nil {
		writeError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, event)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodPost) {
		return
	}
	var req models.Event
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		writeError(w, "bad request", http.StatusBadRequest)
		return
	}

	event, err := h.service.Update(&req)
	if err != nil {
		writeError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, event)
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodPost) {
		return
	}

	idStr := r.URL.Query().Get("event_id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		writeError(w, "invalid event_id", http.StatusBadRequest)
		return
	}

	err = h.service.Delete(uint(id))
	if err != nil {
		writeError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, "event deleted")
}

func (h *Handler) GetForDay(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodGet) {
		return
	}

	userID, err := queryParseUserID(r)
	if err != nil {
		writeError(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	date, err := queryParseDate(r)
	if err != nil {
		writeError(w, "invalid date", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetForDay(userID, date)
	if err != nil {
		writeError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, events)
}

func (h *Handler) GetForWeek(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodGet) {
		return
	}

	userID, err := queryParseUserID(r)
	if err != nil {
		writeError(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	date, err := queryParseDate(r)
	if err != nil {
		writeError(w, "invalid date", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetForWeek(userID, date)
	if err != nil {
		writeError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, events)
}

func (h *Handler) GetForMonth(w http.ResponseWriter, r *http.Request) {
	if !checkMethod(w, r, http.MethodGet) {
		return
	}

	userID, err := queryParseUserID(r)
	if err != nil {
		writeError(w, "invalid user_id", http.StatusBadRequest)
		return
	}
	date, err := queryParseDate(r)
	if err != nil {
		writeError(w, "invalid date", http.StatusBadRequest)
		return
	}

	events, err := h.service.GetForMonth(userID, date)
	if err != nil {
		writeError(w, err.Error(), http.StatusServiceUnavailable)
		return
	}

	writeJSON(w, events)
}
