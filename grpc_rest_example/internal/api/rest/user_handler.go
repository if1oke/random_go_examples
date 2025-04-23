package rest

import (
	"encoding/json"
	"fmt"
	"grpc_rest/internal/domain/entity"
	"net/http"
)

type UserHandler struct{}

func (u *UserHandler) Get(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(entity.User{Username: "Igor Fedorov"})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (u *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("Create user with username: %s\n", user.Username)
}
