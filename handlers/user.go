package handlers

import (
	"github.com/dimasbayuseno/farmacare-test/models"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"net/http"
	"time"
)

func Login(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		secretKey := []byte("your_secret_key")
		expiration := time.Hour

		username, password, ok := c.Request().BasicAuth()
		if !ok {
			return c.String(http.StatusUnauthorized, "Invalid credentials")
		}
		user, err := models.GetUserByUsername(db, username)
		if err != nil {
			return c.String(http.StatusUnauthorized, "Invalid credentials")
		}
		if user.Password != password {
			return c.String(http.StatusUnauthorized, "Invalid credentials")
		}
		token, err := generateToken(user.Role, username, secretKey, expiration)
		if err != nil {
			return c.String(http.StatusInternalServerError, "Failed to generate token")
		}

		return c.JSON(http.StatusOK, map[string]string{
			"token": token,
		})
	}
}

func generateToken(role, username string, secretKey []byte, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"role":     role,
		"username": username,
		"exp":      time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}
