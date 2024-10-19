package auth

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
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
func parseToken(tokenString string) (*jwt.Token, error) {
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

	expiredTime := claims["expired_at"].(time.Time)
	if expiredTime.Unix() > time.Now().Unix() {
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
	return claims["user_id"].(int), nil
}

func VerifyToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "no token provided"})
			return
		}
		if _, err := ValidateJWT(tokenString); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}
		c.Next()
	}
}
