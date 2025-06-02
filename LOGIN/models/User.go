package models

// User representa un usuario en el sistema
type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password,omitempty"` // omitempty para no mostrar en respuestas JSON
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// UserResponse es una versión segura de User para enviar en respuestas
type UserResponse struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

// ToResponse convierte un User a UserResponse (sin password)
func (u *User) ToResponse() UserResponse {
	return UserResponse{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Role:     u.Role,
	}
}

// LoginRequest representa los datos necesarios para iniciar sesión
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterRequest representa los datos necesarios para registrarse
type RegisterRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// UpdateUserRequest representa los datos para actualizar un usuario
type UpdateUserRequest struct {
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}