//go:generate mockgen -package server -source=server.go -destination server_mock.go

package server

import (
	"context"
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
