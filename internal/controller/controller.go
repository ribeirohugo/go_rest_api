//go:generate mockgen -package controller -source=controller.go -destination controller_mock.go

package controller

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

type Controller struct {
	mux     *mux.Router
	service Service
}

func New(service Service) *Controller {
	s := &Controller{
		mux:     mux.NewRouter(),
		service: service,
	}

	return s
}

func (c *Controller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.routing()

	c.mux.ServeHTTP(w, r)
}

func (c *Controller) Mux() *mux.Router {
	return c.mux
}
