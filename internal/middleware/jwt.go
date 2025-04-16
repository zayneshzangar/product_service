package middleware

import (
	"context"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type contextKey string

const (
	ContextUserID contextKey = "user_id"
	ContextRole   contextKey = "role"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// JWTAuthMiddleware проверяет JWT из Authorization или Cookie
func JWTAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var tokenString string

		// 1. Пытаемся взять из заголовка Authorization
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Bearer ") {
			tokenString = strings.TrimPrefix(authHeader, "Bearer ")
		} else {
			// 2. Иначе пробуем взять из cookie
			cookie, err := r.Cookie("token")
			if err != nil || cookie.Value == "" {
				http.Error(w, "missing token", http.StatusUnauthorized)
				return
			}
			tokenString = cookie.Value
		}

		// 3. Парсим токен
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}

		// 4. Извлекаем user_id и role
		userIDFloat, ok1 := claims["user_id"].(float64)
		role, ok2 := claims["role"].(string)
		if !ok1 || !ok2 {
			http.Error(w, "invalid token claims", http.StatusUnauthorized)
			return
		}
		userID := int64(userIDFloat)

		// 5. Добавляем в context
		ctx := context.WithValue(r.Context(), ContextUserID, userID)
		ctx = context.WithValue(ctx, ContextRole, role)

		// 6. Передаём дальше
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// AdminOnlyMiddleware разрешает доступ только admin-пользователям
func AdminOnlyMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		role, ok := r.Context().Value(ContextRole).(string)
		if !ok || role != "admin" {
			http.Error(w, "admin access only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// GetUserIDFromContext возвращает user_id из контекста
func GetUserIDFromContext(ctx context.Context) (int64, bool) {
	userID, ok := ctx.Value(ContextUserID).(int64)
	return userID, ok
}

// GetRoleFromContext возвращает роль из контекста
func GetRoleFromContext(ctx context.Context) (string, bool) {
	role, ok := ctx.Value(ContextRole).(string)
	return role, ok
}
