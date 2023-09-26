package auth

import (
	"encoding/json"
	"net/http"

	"github.com/Samuel200018/pills_backend/auth/application"
	"github.com/Samuel200018/pills_backend/auth/domain"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	authUseCases application.AuthUseCase
}

func NewAuthHandler(authUsesCase application.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUsesCase}
}

var secretKey = []byte("medicine")

func (userHand *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)

	user := domain.User{
		FirstName: r.FormValue("name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Password:  string(hashedPassword),
	}

	createUser, err := userHand.authUseCases.CreateUser(user)
	if err != nil {
		http.Error(w, "Error creating user "+err.Error(), http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "User created successfully",
		"user":    createUser,
	}

	json.NewEncoder(w).Encode(response)

}
