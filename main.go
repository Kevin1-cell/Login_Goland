package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"Login/LOGIN/config"
	"Login/LOGIN/routes"
)

func main() {
	// Inicializar router
	r := mux.NewRouter()

	// Configurar rutas
	routes.SetupRoutes(r)

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	// Inicializar datos de ejemplo
	config.InitSampleData()

	// Iniciar servidor
	port := ":8080"
	log.Printf("Servidor iniciado en http://localhost%s", port)
	log.Fatal(http.ListenAndServe(port, handler))
}