package common

import "github.com/golang-jwt/jwt/v4"

var SecretKey = "M*!7VjT9Dm0pkv2g@dHl"

func GenerateToken(claims *jwt.MapClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	webtoken, err := token.SignedString([]byte(SecretKey))
	if err != nil {
		return "", err
	}

	return webtoken, nil
}
