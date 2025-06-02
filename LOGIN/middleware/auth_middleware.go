package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"Login/LOGIN/config"
	"Login/LOGIN/handlers"
)

// AuthMiddleware verifica que el token JWT sea válido
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener token del encabezado Authorization
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Se requiere autorización")
			return
		}

		// El formato debe ser "Bearer {token}"
		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Formato de autorización inválido")
			return
		}

		// Validar token
		tokenString := tokenParts[1]
		claims := jwt.MapClaims{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return config.JWTSecret, nil
		})

		if err != nil || !token.Valid {
			handlers.RespondWithError(w, http.StatusUnauthorized, "Token inválido")
			return
		}

		// Añadir claims al contexto para uso posterior
		ctx := context.WithValue(r.Context(), "claims", claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminMiddleware verifica que el usuario tenga rol de administrador
func AdminMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Obtener claims del contexto
		claims, ok := r.Context().Value("claims").(jwt.MapClaims)
		if !ok {
			handlers.RespondWithError(w, http.StatusUnauthorized, "No se pudo obtener información del usuario")
			return
		}

		// Verificar rol
		role, ok := claims["role"].(string)
		if !ok || role != "admin" {
			handlers.RespondWithError(w, http.StatusForbidden, "Se requieren permisos de administrador")
			return
		}

		next.ServeHTTP(w, r)
	})
}