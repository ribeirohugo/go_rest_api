package controller

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

func (c *Controller) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	user, err := c.service.FindUser(context.Background(), userID)
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

func (c *Controller) NewUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	returnUser, err := c.service.CreateUser(context.Background(), user)
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

func (c *Controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

	updatedUser, err := c.service.UpdateUser(context.Background(), user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", jsonContentType)

	err = json.NewEncoder(w).Encode(updatedUser)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (c *Controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]

	err := c.service.DeleteUser(context.Background(), userID)
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
