package middleware

import (
	"fmt"
	"github.com/dimasbayuseno/farmacare-test/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
)

func AuthMiddleware(db *gorm.DB, roles ...string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			secretKey := []byte("your_secret_key")

			token := c.Request().Header.Get("Authorization")
			if token == "" {
				return c.String(http.StatusUnauthorized, "Missing token")
			}

			claims, err := ParseToken(token, secretKey)
			if err != nil {
				return c.String(http.StatusUnauthorized, "Invalid token")
			}

			username, ok := claims["username"].(string)
			if !ok {
				return c.String(http.StatusUnauthorized, "Invalid token")
			}

			user, err := models.GetUserByUsername(db, username)
			if err != nil {
				return c.String(http.StatusInternalServerError, "Failed to retrieve user")
			}

			if !hasRole(user.Role, roles) {
				return c.String(http.StatusForbidden, "Access denied")
			}

			c.Set("user", user)

			return next(c)
		}
	}
}

func hasRole(userRole string, roles []string) bool {
	for _, role := range roles {
		if userRole == role {
			return true
		}
	}
	return false
}

func ParseToken(tokenString string, secretKey []byte) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
