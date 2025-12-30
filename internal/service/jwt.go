package auth

import (
	"context"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func (s *AuthService) generateJWTTokens(
	ctx context.Context,
	userId string,
	dormitoryId int,
) (string, string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":      userId,
		"dormitory_id": dormitoryId,
		"exp":          time.Now().Add(15 * time.Minute).Unix(),
		"iat":          time.Now().Unix(),
		"type":         "access",
	})

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userId,
		"exp":     time.Now().Add(7 * 24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
		"type":    "refresh",
	})

	accessTokenString, err := accessToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", "", fmt.Errorf("%w: error while generating access token: %v", ErrInternal, err)
	}

	refreshTokenString, err := refreshToken.SignedString([]byte(s.jwtSecret))
	if err != nil {
		return "", "", fmt.Errorf("%w: error while generating refresh token: %v", ErrInternal, err)
	}

	return accessTokenString, refreshTokenString, nil
}
