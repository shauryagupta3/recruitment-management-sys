package handlers

import (
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt"
)

func AdminProtectedHandler(w http.ResponseWriter, r *http.Request) (jwt.MapClaims, error) {
	w.Header().Set("Content-Type", "application/json")
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		return nil, NewAPIError(http.StatusUnauthorized, "missing auth")
	}
	tokenString = tokenString[len("Bearer "):]

	claims, err := verifyToken(tokenString)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Invalid token")
		return nil, err
	}

	userType, err := getTypeFromClaims(claims)
	if err != nil {
		return nil, err
	}

	if userType != "admin" {
		w.WriteHeader(http.StatusForbidden)
		return nil, NewAPIError(http.StatusUnauthorized, "unauthorized")
	}

	return claims,nil
}

func verifyToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return SECRET, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, err
	}
	return nil, fmt.Errorf("error")
}

func getIDFromClaims(claims jwt.MapClaims) (string, error) {
	if id, ok := claims["id"].(string); ok {
		return id, nil
	}
	return "", fmt.Errorf("id not found in token claims")
}

func getTypeFromClaims(claims jwt.MapClaims) (string, error) {
	if UserType, ok := claims["type"].(string); ok {
		return UserType, nil
	}
	return "", fmt.Errorf("type not found in token claims")
}
