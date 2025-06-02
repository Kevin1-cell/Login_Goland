package models

// Response es una estructura genérica para respuestas HTTP
type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// TokenResponse es la respuesta cuando se genera un token JWT
type TokenResponse struct {
	Token string `json:"token"`
}