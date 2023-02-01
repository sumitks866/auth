package auth

import (
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/sumitks866/auth/models"
)

type Claims struct {
	Username string `json:"username"`
	UserId   string `json:"user_id"`
	jwt.RegisteredClaims
}

var jwtSecret []byte

func Setup() {
	key, exists := os.LookupEnv("JWT_SECRET")

	if !exists {
		log.Fatal("JWT_SECRET not present")
	}

	jwtSecret = []byte(key)
}

// Generates a new JWT Token for a user
func GenerateToken(user models.User) (string, error) {
	currTime := time.Now()
	expire := currTime.Add(20 * time.Second)

	claims := Claims{
		Username: user.Username,
		UserId:   user.ID.Hex(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expire),
			Issuer:    "foodile-auth",
		},
	}

	tokenClaim := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	token, err := tokenClaim.SignedString(jwtSecret)

	return token, err
}

// Parses JWT token to claims
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, nil
	}

	return nil, err
}
