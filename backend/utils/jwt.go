package utils

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtDuration = time.Hour * 24
	jwtSecret   = "jwtSecret"
)

type MyClaims struct {
	UserID     uint `json:"userid"`
	UserTypeID uint `json:"usertype"`
	jwt.RegisteredClaims
}

func GenerateJWT(userID int, UserTypeID uint) (string, error) {
	// Setear expiracion
	expirationTime := time.Now().Add(jwtDuration)

	// Construir los claims
	claims := MyClaims{
		UserID:     uint(userID),
		UserTypeID: UserTypeID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "backend",
			Subject:   "auth",
		},
	}
	// Crear el token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Firmar el token
	tokenString, err := token.SignedString([]byte(jwtSecret))
	if err != nil {
		return "", fmt.Errorf("error generating token: %w", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secret := jwtSecret
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check claims, e.g., expiration time
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return false, fmt.Errorf("token has expired")
			}
		}
		// you can check other claims here if needed
		return true, nil
	} else {
		return false, fmt.Errorf("invalid token")
	}
}
