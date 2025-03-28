package jwt

import (
	"fmt"
	"minisapi/services/auth/internal/configs"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTManager struct {
	secretKey       []byte
	accessTokenTTL  time.Duration
	refreshTokenTTL time.Duration
}

func NewJWTManager(cfg configs.JWTConfig) (*JWTManager, error) {
	expiration, err := time.ParseDuration(cfg.Expiration)
	if err != nil {
		return nil, fmt.Errorf("invalid expiration duration: %v", err)
	}

	return &JWTManager{
		secretKey:       []byte(cfg.Secret),
		accessTokenTTL:  expiration,
		refreshTokenTTL: expiration * 2,
	}, nil
}

func (m *JWTManager) GenerateAccessToken(userID uint) (string, error) {
	return m.generateToken(userID, m.accessTokenTTL)
}

func (m *JWTManager) GenerateRefreshToken(userID uint) (string, error) {
	return m.generateToken(userID, m.refreshTokenTTL)
}

func (m *JWTManager) GenerateResetToken(userID uint) (string, error) {
	return m.generateToken(userID, 1*time.Hour)
}

func (m *JWTManager) GenerateVerifyToken(userID uint) (string, error) {
	return m.generateToken(userID, 24*time.Hour)
}

func (m *JWTManager) generateToken(userID uint, expiration time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(expiration).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(m.secretKey)
}

func (m *JWTManager) ValidateToken(tokenString string) (uint, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return m.secretKey, nil
	})

	if err != nil {
		return 0, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return 0, fmt.Errorf("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid user_id in token")
	}

	return uint(userID), nil
}
