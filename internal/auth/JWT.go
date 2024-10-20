package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

var secretKey = []byte("mysecretkey")

func GenerateJWT(userID int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":    userID,
		"expired_at": time.Now().Add(time.Hour * 1).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func parseToken(authHeader string) (*jwt.Token, error) {
	tokenString := strings.TrimPrefix(authHeader, "Bearer ") // удаление префикса "Bearer "
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	return jwtToken, nil
}
func getClaims(jwtToken *jwt.Token) (jwt.MapClaims, error) {
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid token")
	}
	return claims, nil
}
func ValidateJWT(tokenString string) (bool, error) {
	jwtToken, err := parseToken(tokenString)
	if err != nil {
		return false, err
	}
	claims, err := getClaims(jwtToken)

	expiredTime := claims["expired_at"].(float64)
	now := float64(time.Now().Unix())
	if expiredTime <= now {
		return false, fmt.Errorf("token is expired")
	}
	return true, nil
}

func GetUserIdFromJwt(token string) (int, error) {
	jwtToken, err := parseToken(token)
	if err != nil {
		return 0, err
	}

	claims, err := getClaims(jwtToken)
	if err != nil {
		return 0, err
	}

	// Преобразуем значение user_id из float64 в int
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, fmt.Errorf("invalid token: user_id is not a number")
	}

	return int(userID), nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.Request.Header.Get("Authorization") // получение заголовка Authorization

		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			c.Abort()
			return
		}
		if _, err := ValidateJWT(authHeader); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			c.Abort()
			return
		}
		c.Next()
	}
}
