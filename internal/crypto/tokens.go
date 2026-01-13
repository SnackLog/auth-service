package crypto

import (
	"fmt"
	"time"

	"github.com/SnackLog/auth-service/internal/config"
	serviceconfiglib "github.com/SnackLog/service-config-lib"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func CreateAuthToken(username string) (string, error) {
	signKey := config.GetConfig().JwtSignKey
	claims := jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
		IssuedAt:  jwt.NewNumericDate(time.Now()),
		NotBefore: jwt.NewNumericDate(time.Now()),
		Issuer:    serviceconfiglib.GetConfig().ServiceName,
		Subject:   username,
		ID:        uuid.New().String(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(signKey))
	if err != nil {
		return "", fmt.Errorf("Failed to cryptographically sign token: %v", err)
	}
	return signedToken, nil
}

func ParseAndValidateToken(tokenString string) (*jwt.RegisteredClaims, error) {
	signKey := config.GetConfig().JwtSignKey
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(signKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	switch {
	case err != nil:
		return nil, fmt.Errorf("Failed to parse token")
	case !token.Valid:
		return nil, fmt.Errorf("Token is invalid")
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return nil, fmt.Errorf("Failed to parse token claims")
	}
	return claims, nil
}
