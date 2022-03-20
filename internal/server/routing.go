package server

import "net/http"

func (s *server) routing() {
	s.mux.HandleFunc("/user/{id}", s.GetUser).Methods(http.MethodGet)
	s.mux.HandleFunc("/users", s.NewUser).Methods(http.MethodPost)
	s.mux.HandleFunc("/user", s.UpdateUser).Methods(http.MethodPut, http.MethodPatch)
	s.mux.HandleFunc("/user/{id}", s.DeleteUser).Methods(http.MethodDelete)
}
