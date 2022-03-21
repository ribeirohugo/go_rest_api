package controller

import "net/http"

func (c *Controller) routing() {
	c.mux.HandleFunc("/user/{id}", c.GetUser).Methods(http.MethodGet)
	c.mux.HandleFunc("/users", c.NewUser).Methods(http.MethodPost)
	c.mux.HandleFunc("/user", c.UpdateUser).Methods(http.MethodPut, http.MethodPatch)
	c.mux.HandleFunc("/user/{id}", c.DeleteUser).Methods(http.MethodDelete)
}
