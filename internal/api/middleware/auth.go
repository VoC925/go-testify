package middleware

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

const (
	// Объявляем привилегии нашей системы
	UpdatePermission = "update"

	// Объявляем роли нашей системы
	AdminRole = "admin"
)

var (
	// Связка роль — привилегии
	rolePermissions = map[string][]string{
		AdminRole: {UpdatePermission},
	}

	// Связка пользователь — роль
	adminRoles = map[string][]string{
		"admin": {AdminRole},
	}

	secretKey = []byte("gBElG5NThZSye")
)

// аутентификация
func verifyUser(token string, permission string) bool {
	jwtToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		fmt.Printf("Failed to parse token: %s\n", err)
		return false
	}
	if !jwtToken.Valid {
		return false
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return false
	}
	loginRaw, ok := claims["login"]
	if !ok {
		return false
	}

	login, ok := loginRaw.(string)
	if !ok {
		return false
	}

	for _, roles := range adminRoles[login] {
		for _, storedPermission := range rolePermissions[roles] {
			if permission == storedPermission {
				return true
			}
		}
	}
	return false
}

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
