package config

import (
	"time"

	"Login/LOGIN/models"
)

// Clave secreta para JWT
var JWTSecret = []byte("mi_clave_super_secreta")

// Duración del token JWT
var TokenDuration = time.Hour * 24

// Almacenamiento temporal de usuarios
var Users = []models.User{}

// Inicializar datos de ejemplo
func InitSampleData() {
	// Crear algunos usuarios de ejemplo
	Users = append(Users, models.User{
		ID:       "1",
		Username: "admin",
		Password: "admin123", // En producción, esto debería estar hasheado
		Email:    "admin@example.com",
		Role:     "admin",
	})

	Users = append(Users, models.User{
		ID:       "2",
		Username: "usuario",
		Password: "user123", // En producción, esto debería estar hasheado
		Email:    "usuario@example.com",
		Role:     "user",
	})
}

// Función para encontrar un usuario por ID
func FindUserByID(id string) *models.User {
	for i, user := range Users {
		if user.ID == id {
			return &Users[i]
		}
	}
	return nil
}

// Función para encontrar un usuario por nombre de usuario
func FindUserByUsername(username string) *models.User {
	for i, user := range Users {
		if user.Username == username {
			return &Users[i]
		}
	}
	return nil
}

// Función para eliminar un usuario por ID
func DeleteUserByID(id string) bool {
	for i, user := range Users {
		if user.ID == id {
			// Eliminar usuario del slice
			Users = append(Users[:i], Users[i+1:]...)
			return true
		}
	}
	return false
}

// Función para actualizar un usuario
func UpdateUser(updatedUser models.User) bool {
	for i, user := range Users {
		if user.ID == updatedUser.ID {
			// Actualizar usuario
			Users[i] = updatedUser
			return true
		}
	}
	return false
}