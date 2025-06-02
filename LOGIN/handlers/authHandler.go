package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"Login/LOGIN/config"
	"Login/LOGIN/models"
)

// LoginHandler maneja la autenticación de usuarios
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var loginReq models.LoginRequest
	
	// Decodificar el cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&loginReq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Formato de solicitud inválido")
		return
	}

	// Buscar usuario por nombre de usuario
	user := config.FindUserByUsername(loginReq.Username)
	if user == nil || user.Password != loginReq.Password {
		respondWithError(w, http.StatusUnauthorized, "Credenciales inválidas")
		return
	}

	// Generar token JWT
	token, err := generateJWT(user)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error al generar token")
		return
	}

	// Responder con el token
	respondWithJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Login exitoso",
		Data: models.TokenResponse{
			Token: token,
		},
	})
}

// RegisterHandler maneja el registro de nuevos usuarios
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var registerReq models.RegisterRequest
	
	// Decodificar el cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&registerReq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Formato de solicitud inválido")
		return
	}

	// Verificar si el usuario ya existe
	existingUser := config.FindUserByUsername(registerReq.Username)
	if existingUser != nil {
		respondWithError(w, http.StatusConflict, "El nombre de usuario ya está en uso")
		return
	}

	// Crear nuevo usuario
	newID := strconv.Itoa(len(config.Users) + 1)
	newUser := models.User{
		ID:       newID,
		Username: registerReq.Username,
		Password: registerReq.Password, // En producción, esto debería estar hasheado
		Email:    registerReq.Email,
		Role:     "user", // Rol por defecto
	}

	// Guardar usuario
	config.Users = append(config.Users, newUser)

	// Generar token JWT
	token, err := generateJWT(&newUser)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Error al generar token")
		return
	}

	// Responder con el token
	respondWithJSON(w, http.StatusCreated, models.Response{
		Success: true,
		Message: "Registro exitoso",
		Data: models.TokenResponse{
			Token: token,
		},
	})
}

// generateJWT genera un token JWT para un usuario
func generateJWT(user *models.User) (string, error) {
	// Crear claims
	claims := jwt.MapClaims{
		"id":       user.ID,
		"username": user.Username,
		"role":     user.Role,
		"exp":      time.Now().Add(config.TokenDuration).Unix(),
	}

	// Crear token con claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar token con clave secreta
	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// respondWithJSON envía una respuesta JSON
func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// respondWithError envía una respuesta de error
func respondWithError(w http.ResponseWriter, code int, message string) {
	respondWithJSON(w, code, models.Response{
		Success: false,
		Message: message,
	})
}