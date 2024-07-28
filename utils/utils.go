package utils

import (
	"encoding/base64"
	"fmt"
	"strings"

	"github.com/dgrijalva/jwt-go"

	"karma_files_go/config"
)

// Function to decode the token
func Decode(kp string) (map[string]interface{}, error) {
	// Remove the first 3 characters
	if len(kp) <= 3 {
		return nil, fmt.Errorf("input string is too short")
	}
	base := kp[3:]

	// Add padding if necessary
	if len(base)%4 != 0 {
		base += strings.Repeat("=", 4-len(base)%4)
	}

	// Decode the base64 string
	decodedBytes, err := base64.StdEncoding.DecodeString(base)
	if err != nil {
		return nil, fmt.Errorf("base64 decoding failed: %w", err)
	}
	tokenString := string(decodedBytes)

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
