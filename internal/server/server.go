//go:generate mockgen -package server -source=server.go -destination server_mock.go

package server

import (
	"context"
	"net/http"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/gorilla/mux"
)

// UserService abstracts the service layer.
type UserService interface {
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

type server struct {
	mux     *mux.Router
	service UserService
}

func New(service UserService) *server {
	s := &server{
		mux:     mux.NewRouter(),
		service: service,
	}

	s.mux.HandleFunc("/users", s.GetUser).Methods(http.MethodGet)
	s.mux.HandleFunc("/users", s.NewUser).Methods(http.MethodPost)
	s.mux.HandleFunc("/users", s.UpdateUser).Methods(http.MethodPut, http.MethodPatch)
	s.mux.HandleFunc("/users", s.DeleteUser).Methods(http.MethodDelete)

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.mux.ServeHTTP(w, r)
}
