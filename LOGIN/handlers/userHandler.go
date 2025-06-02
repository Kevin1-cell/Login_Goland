package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"Login/LOGIN/config"
	"Login/LOGIN/models"
)

// GetAllUsersHandler obtiene todos los usuarios
func GetAllUsersHandler(w http.ResponseWriter, r *http.Request) {
	// Convertir usuarios a respuesta segura (sin passwords)
	var usersResponse []models.UserResponse
	for _, user := range config.Users {
		usersResponse = append(usersResponse, user.ToResponse())
	}

	respondWithJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Usuarios obtenidos con éxito",
		Data:    usersResponse,
	})
}

// GetUserByIDHandler obtiene un usuario por su ID
func GetUserByIDHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener ID de los parámetros de la URL
	params := mux.Vars(r)
	id := params["id"]

	// Buscar usuario
	user := config.FindUserByID(id)
	if user == nil {
		respondWithError(w, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	// Responder con usuario (sin password)
	respondWithJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Usuario obtenido con éxito",
		Data:    user.ToResponse(),
	})
}

// CreateUserHandler crea un nuevo usuario (solo para administradores)
func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	var newUser models.User
	
	// Decodificar el cuerpo de la solicitud
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Formato de solicitud inválido")
		return
	}

	// Verificar si el usuario ya existe
	existingUser := config.FindUserByUsername(newUser.Username)
	if existingUser != nil {
		respondWithError(w, http.StatusConflict, "El nombre de usuario ya está en uso")
		return
	}

	// Asignar ID
	newUser.ID = strconv.Itoa(len(config.Users) + 1)

	// Guardar usuario
	config.Users = append(config.Users, newUser)

	// Responder con usuario creado (sin password)
	respondWithJSON(w, http.StatusCreated, models.Response{
		Success: true,
		Message: "Usuario creado con éxito",
		Data:    newUser.ToResponse(),
	})
}

// UpdateUserHandler actualiza un usuario existente
func UpdateUserHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener ID de los parámetros de la URL
	params := mux.Vars(r)
	id := params["id"]

	// Buscar usuario
	user := config.FindUserByID(id)
	if user == nil {
		respondWithError(w, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	// Decodificar datos de actualización
	var updateReq models.UpdateUserRequest
	err := json.NewDecoder(r.Body).Decode(&updateReq)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Formato de solicitud inválido")
		return
	}

	// Actualizar campos si están presentes
	if updateReq.Username != "" {
		// Verificar si el nuevo nombre de usuario ya está en uso
		if updateReq.Username != user.Username {
			existingUser := config.FindUserByUsername(updateReq.Username)
			if existingUser != nil {
				respondWithError(w, http.StatusConflict, "El nombre de usuario ya está en uso")
				return
			}
		}
		user.Username = updateReq.Username
	}

	if updateReq.Email != "" {
		user.Email = updateReq.Email
	}

	if updateReq.Role != "" {
		user.Role = updateReq.Role
	}

	// Guardar cambios
	config.UpdateUser(*user)

	// Responder con usuario actualizado
	respondWithJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Usuario actualizado con éxito",
		Data:    user.ToResponse(),
	})
}

// DeleteUserHandler elimina un usuario
func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	// Obtener ID de los parámetros de la URL
	params := mux.Vars(r)
	id := params["id"]

	// Eliminar usuario
	success := config.DeleteUserByID(id)
	if !success {
		respondWithError(w, http.StatusNotFound, "Usuario no encontrado")
		return
	}

	// Responder con éxito
	respondWithJSON(w, http.StatusOK, models.Response{
		Success: true,
		Message: "Usuario eliminado con éxito",
	})
}