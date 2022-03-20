//go:generate mockgen -package server -source=server.go -destination server_mock.go

package server

import (
	"context"
	"net/http"

	"github.com/ribeirohugo/golang_startup/internal/model"

	"github.com/gorilla/mux"
)

// Service abstracts the service layer.
type Service interface {
	FindUser(ctx context.Context, id string) (model.User, error)
	CreateUser(ctx context.Context, user model.User) (model.User, error)
	UpdateUser(ctx context.Context, user model.User) (model.User, error)
	DeleteUser(ctx context.Context, id string) error
}

type server struct {
	mux     *mux.Router
	service Service
}

func New(service Service) *server {
	s := &server{
		mux:     mux.NewRouter(),
		service: service,
	}

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.routing()

	s.mux.ServeHTTP(w, r)
}
