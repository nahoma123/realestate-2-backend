package user

import (
	"errors"
	"fmt"
	"time"
	"visitor_management/internal/constant"
	"visitor_management/internal/constant/model"

	"github.com/golang-jwt/jwt/v4"
)

func (o *user) VerifyToken(signingMethod jwt.SigningMethod, tokenString string) (bool, *jwt.MapClaims, error) {
	claims := &jwt.MapClaims{}

	secretKey := []byte(constant.GetConfig().JWT_SECRET)

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if t.Method.Alg() != signingMethod.Alg() {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return false, claims, err
	}
	if !token.Valid {
		return false, claims, errors.New("invalid token")
	}
	return true, claims, nil
}

func (o *user) GenerateJWT(user *model.User) (string, error) {
	// The secret key for signing the token
	secretKey := []byte(constant.GetConfig().JWT_SECRET)

	claims := jwt.MapClaims{
		"user_id": user.UserID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}
