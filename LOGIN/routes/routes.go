package routes

import (
	"net/http"

	"github.com/gorilla/mux"

	"Login/LOGIN/handlers"
	"Login/LOGIN/middleware"
)

// SetupRoutes configura todas las rutas de la API
func SetupRoutes(r *mux.Router) {
	// Crear subrouter para la API
	api := r.PathPrefix("/api").Subrouter()

	// Rutas públicas (sin autenticación)
	api.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	api.HandleFunc("/auth/register", handlers.RegisterHandler).Methods("POST")
	
	// Rutas protegidas (requieren autenticación)
	protected := api.PathPrefix("/").Subrouter()
	protected.Use(middleware.AuthMiddleware)

	// Rutas de usuarios (accesibles para todos los usuarios autenticados)
	protected.HandleFunc("/users", handlers.GetAllUsersHandler).Methods("GET")
	protected.HandleFunc("/users/{id}", handlers.GetUserByIDHandler).Methods("GET")
	
	// Rutas de administración (solo para administradores)
	admin := protected.PathPrefix("/admin").Subrouter()
	admin.Use(middleware.AdminMiddleware)
	
	admin.HandleFunc("/users", handlers.CreateUserHandler).Methods("POST")
	admin.HandleFunc("/users/{id}", handlers.UpdateUserHandler).Methods("PUT")
	admin.HandleFunc("/users/{id}", handlers.DeleteUserHandler).Methods("DELETE")

	// Ruta para verificar que el servidor está funcionando
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("API funcionando correctamente"))
	}).Methods("GET")
}