package server

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/gorilla/mux"
)

const (
	jsonContentType = "application/json"

	userDeletedMessage = "user successfully removed"
)

func (s *server) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := s.service.GetUserByEmail(context.Background(), userID)
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
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	returnUser, err := s.service.CreateUser(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(returnUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *server) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	err = s.service.UpdateUser(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
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
	vars := mux.Vars(r)
	userID := vars["id"]

	err := s.service.DeleteUser(context.Background(), userID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(userDeletedMessage)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
