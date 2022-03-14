package server

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

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
