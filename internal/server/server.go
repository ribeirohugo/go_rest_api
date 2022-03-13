package server

import (
	"bytes"
	"context"
	"fmt"
	"net/http"

	"github.com/ribeirohugo/golang_startup/internal/model"
)

// UserService abstracts the service layer.
type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type server struct {
	mux     *http.ServeMux
	service UserService
}

func New(service UserService) *server {
	s := &server{
		mux:     http.NewServeMux(),
		service: service,
	}

	s.mux.HandleFunc("/users/email", s.GetSingleUserByEmail)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}

func (s *server) GetSingleUserByEmail(w http.ResponseWriter, r *http.Request) {
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

	output := fmt.Sprintf("User with %s email is named %s.", user.Email, user.Name)

	resp := bytes.NewBufferString(output)

	_, _ = w.Write(resp.Bytes())
}
