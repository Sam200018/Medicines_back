package auth

import (
	"encoding/json"
	"github.com/golang-jwt/jwt"
	"net/http"
	"time"

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

func (authHand *AuthHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(r.FormValue("password")), bcrypt.DefaultCost)

	user := domain.User{
		FirstName: r.FormValue("name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Password:  string(hashedPassword),
	}

	createUser, err := authHand.authUseCases.CreateUser(user)
	if err != nil {
		response := map[string]interface{}{
			"message": "Error creating user",
		}
		w.WriteHeader(http.StatusConflict)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := map[string]interface{}{
		"message": "User created successfully",
		"user":    createUser,
	}

	json.NewEncoder(w).Encode(response)

}

func (authHand *AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	user, err := authHand.authUseCases.GetUser(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusBadRequest)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(r.FormValue("password")))

	if err != nil {
		http.Error(w, "Invalid password", http.StatusBadRequest)
		return
	}

	tokenString, err := newToken(user)

	if err != nil {
		http.Error(w, "Error creating token", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message": "Login successfully",
		"token":   tokenString,
	}

	json.NewEncoder(w).Encode(response)
}

func newToken(user domain.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":      user.Email,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"exp":        time.Now().Add(time.Hour * 72).Unix(),
	})

	signedString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedString, err
}

func (authHand *AuthHandler) Logout(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"message": "Logout successfully",
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func (authHand *AuthHandler) CheckStatus(w http.ResponseWriter, r *http.Request) {
	email := r.FormValue("email")
	user, err := authHand.authUseCases.GetUser(email)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if user.ID == 0 {
		http.Error(w, "User not found", http.StatusBadRequest)

	}

	tokenString, err := newToken(user)

	if err != nil {
		http.Error(w, "Error creating token", http.StatusBadRequest)
		return
	}

	response := map[string]interface{}{
		"message":       "Login successfully",
		"updated_token": tokenString,
	}

	json.NewEncoder(w).Encode(response)
}
