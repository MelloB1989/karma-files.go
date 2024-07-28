package utils

import (
	"fmt"

	"github.com/dgrijalva/jwt-go"

	"karma_files_go/config"
)

// Function to decode the token
func Decode(tokenString string) (map[string]interface{}, error) {
	// Decode the JWT token
	// token, _, err := new(jwt.Parser).ParseUnverified(tokenString, jwt.MapClaims{})
	// if err != nil {
	// 	return nil, fmt.Errorf("jwt parsing failed: %w", err)
	// }

	// Parse the JWT token
	token, err := jwt.ParseWithClaims(
		tokenString,
		jwt.MapClaims{},
		func(token *jwt.Token) (interface{}, error) {
			// Make sure the token's algorithm is what you expect:
			return []byte(config.NewConfig().JWTSecret), nil
		},
	)
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		return claims, nil
	}

	return nil, fmt.Errorf("could not parse token claims")
}
