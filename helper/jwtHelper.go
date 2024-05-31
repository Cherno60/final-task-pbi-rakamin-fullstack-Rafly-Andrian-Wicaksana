package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"goAPI/models"
	"os"
	"time"
)

type JWTClaims struct {
	Uuid string `json:"uuid"`
	jwt.RegisteredClaims
}

func GenerateToken(user models.Users) (string, error) {
	claims := JWTClaims{
		Uuid: user.Uuid,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(60 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func ValidateToken(tokenString string) (*string, error) {
	// Use ParseWithClaims to parse and validate the token
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		// Ensure the signing method is as expected
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrSignatureInvalid) {
			return nil, errors.New("Invalid token signature")
		}
		return nil, errors.New("Invalid token")
	}
	//fmt.Println(token)

	claims, ok := token.Claims.(*JWTClaims)
	if !ok || !token.Valid {
		return nil, errors.New("Invalid token")
	}

	//fmt.Println(claims.Uuid)

	return &claims.Uuid, nil
}
