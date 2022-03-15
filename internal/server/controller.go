package server

import (
	"context"
	"encoding/json"
	"net/http"
)

const (
	jsonContentType = "application/json"
)

func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	emails, exists := r.URL.Query()["email"]
	if !exists || len(emails) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.service.GetUserByEmail(context.Background(), emails[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) NewUser(w http.ResponseWriter, r *http.Request) {
	emails, exists := r.URL.Query()["email"]
	if !exists || len(emails) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.service.GetUserByEmail(context.Background(), emails[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	emails, exists := r.URL.Query()["email"]
	if !exists || len(emails) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.service.GetUserByEmail(context.Background(), emails[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) DeleteUser(w http.ResponseWriter, r *http.Request) {
	emails, exists := r.URL.Query()["email"]
	if !exists || len(emails) < 1 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	user, err := s.service.GetUserByEmail(context.Background(), emails[0])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
